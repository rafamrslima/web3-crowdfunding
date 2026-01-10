package repositories

import (
	"context"
	"encoding/hex"
	"log"
	"strings"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/models"
)

func GetAvailableRefundsByDonor(donor []byte) ([]dtos.RefundViewDTO, error) {
	pool, err := GetDB()
	if err != nil {
		return nil, err
	}

	hexStr := strings.TrimPrefix(string(donor), "0x")
	addrBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	rows, err := pool.Query(ctx,
		`SELECT c.campaign_id, d.donor, c.deadline_ts, c.title, c.description, 
    sum(d.amount) as donor_total, 
    (SELECT sum(d2.amount) FROM donations d2 WHERE d2.campaign_id = c.campaign_id) as total_collected,
    c.target_amount, c.image
    FROM campaigns c 
    inner join donations d on c.campaign_id = d.campaign_id
	left join refunds r on c.campaign_id = r.campaign_id and r.donor = d.donor
    WHERE c.deadline_ts < EXTRACT(EPOCH FROM now())
    AND c.target_amount > 0
    AND d.donor = $1
	AND r.tx_hash IS NULL 
    GROUP BY c.campaign_id, d.donor, c.deadline_ts, c.title, c.description, c.image, c.target_amount
    HAVING (SELECT sum(d2.amount) FROM donations d2 WHERE d2.campaign_id = c.campaign_id) < c.target_amount`, addrBytes)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var results []dtos.RefundViewDTO

	for rows.Next() {
		var res dtos.RefundViewDTO
		if err := rows.Scan(
			&res.Donation.CampaignId, &res.Donation.Donor, &res.CampaignDeadline, &res.Donation.Title, &res.Donation.Description,
			&res.Donation.Amount, &res.CampaignAmountCollected, &res.CampaignTarget, &res.Donation.Image); err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}

func SaveRefundIssued(refund models.RefundDbEntity) error {
	pool, err := GetDB()
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, err = pool.Exec(ctx,
		`INSERT INTO refunds (campaign_id, donor, total_contributed, tx_hash, block_number, block_time) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		refund.CampaignId, refund.Donor, refund.TotalContributed, refund.TxHash, refund.BlockNumber, refund.BlockTime)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Row inserted successfully into refunds.")
	return nil
}
