package repositories

import (
	"context"
	"log"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/models"
)

func SaveDonationReceived(donation models.DonationDbEntity) error {
	pool, err := GetDB()
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, err = pool.Exec(ctx,
		`INSERT INTO donations (campaign_id, donor, amount, tx_hash, block_number, block_time) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		donation.CampaignId, donation.Donor, donation.Amount, donation.TxHash, donation.BlockNumber, donation.BlockTime)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Row inserted successfully into donations.")
	return nil
}

func GetDonationsByDonor(owner []byte) ([]dtos.DonationViewDTO, error) {
	pool, err := GetDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	rows, err := pool.Query(ctx,
		`SELECT d.donor, c.campaign_id, c.title, c.description, d.created_at, c.image, d.amount
		FROM donations d inner join campaigns c on d.campaign_id = c.campaign_id
		WHERE d.donor = $1 ORDER BY d.created_at desc`, owner)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var results []dtos.DonationViewDTO

	for rows.Next() {
		var res dtos.DonationViewDTO
		if err := rows.Scan(
			&res.Donor, &res.CampaignId, &res.Title, &res.Description, &res.CreatedAt, &res.Image, &res.Amount); err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}
