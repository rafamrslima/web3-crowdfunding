package db

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

func connect() (*pgxpool.Pool, error) {
	connString := os.Getenv("DATABASE_CONNECTION_STRING")
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	config.MaxConns = 20
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Error to connect to database.", err)
		return nil, err
	}
	return pool, nil
}

func SaveCampaignDraft(creationId string, address common.Address, title string, description string, image string) error {
	pool, err := connect()
	if err != nil {
		return err
	}
	defer pool.Close()

	ctx := context.Background()

	_, err = pool.Exec(ctx,
		`INSERT INTO campaign_drafts (creation_id, owner, title, description, image) 
		VALUES ($1, $2, $3, $4, $5)`,
		creationId, address, title, description, image)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Row inserted successfully into campaign_drafts.")
	return nil
}

func SaveCampaignCreated(campaign models.CampaignDbEntity) error {
	pool, err := connect()
	if err != nil {
		return err
	}
	defer pool.Close()

	ctx := context.Background()

	_, err = pool.Exec(ctx,
		`INSERT INTO campaigns 
		(campaign_id, owner, title, description, target_amount, deadline_ts, image, tx_hash, block_number, block_time) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		campaign.Id,
		campaign.Owner,
		campaign.Title,
		campaign.Description,
		campaign.Target,
		campaign.Deadline,
		campaign.Image,
		campaign.TxHash,
		campaign.BlockNumber,
		campaign.BlockTime)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Row inserted successfully into campaigns.")
	return nil
}

func SaveDonationReceived(donation models.DonationDbEntity) error {
	pool, err := connect()
	if err != nil {
		return err
	}
	defer pool.Close()

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

func SaveRefundIssued(refund models.RefundDbEntity) error {
	pool, err := connect()
	if err != nil {
		return err
	}
	defer pool.Close()

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

func SaveWithdrawCompletion(withdraw models.WithdrawDbEntity) error {
	pool, err := connect()
	if err != nil {
		return err
	}
	defer pool.Close()

	ctx := context.Background()

	_, err = pool.Exec(ctx,
		`INSERT INTO withdrawals (campaign_id, owner, amount, tx_hash, block_number, block_time) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		withdraw.CampaignId, withdraw.Owner, withdraw.Amount, withdraw.TxHash, withdraw.BlockNumber, withdraw.BlockTime)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Row inserted successfully into withdrawals.")
	return nil
}

func FetchAllCampaigns() ([]dtos.CampaignDto, error) {
	pool, err := connect()
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	ctx := context.Background()
	rows, err := pool.Query(ctx,
		`SELECT c.owner, c.title, c.description, c.target_amount, c.deadline_ts, c.image, sum(d.amount) as amount_collected
		FROM campaigns c left join donations d on c.campaign_id = d.campaign_id 
	    GROUP BY c.campaign_id, c.owner, c.title, c.description, c.target_amount, c.deadline_ts, c.image 
		ORDER BY c.campaign_id 
		LIMIT 100`)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var results []dtos.CampaignDto

	for rows.Next() {
		var res dtos.CampaignDto
		if err := rows.Scan(&res.Owner, &res.Title, &res.Description, &res.Target, &res.Deadline, &res.Image, &res.AmountCollected); err != nil {
			log.Println(err)
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}

func GetCampaignsByOwner(owner []byte) ([]dtos.CampaignDto, error) {
	pool, err := connect()
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	ctx := context.Background()
	rows, err := pool.Query(ctx,
		`SELECT c.owner, c.title, c.description, c.target_amount, c.deadline_ts, c.image, sum(d.amount) as amount_collected
		FROM campaigns c left join donations d on c.campaign_id = d.campaign_id 
	    WHERE c.owner = $1 GROUP BY c.owner, c.title, c.description, c.target_amount, c.deadline_ts, c.image
		LIMIT 100`, owner)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var results []dtos.CampaignDto

	for rows.Next() {
		var res dtos.CampaignDto
		if err := rows.Scan(&res.Owner, &res.Title, &res.Description, &res.Target, &res.Deadline, &res.Image, &res.AmountCollected); err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}

func GetCampaignMetadataFromDraft(owner common.Address, creationId string) (models.CampaignMetadata, error) {
	pool, err := connect()
	if err != nil {
		return models.CampaignMetadata{}, err
	}
	defer pool.Close()

	ctx := context.Background()
	rows, err := pool.Query(ctx,
		`SELECT title, description, image FROM campaign_drafts WHERE owner = $1 AND creation_id = $2 LIMIT 1`, owner, creationId)

	if err != nil {
		return models.CampaignMetadata{}, err
	}
	defer rows.Close()

	var campaignMetadata models.CampaignMetadata
	for rows.Next() {
		if err := rows.Scan(&campaignMetadata.Title, &campaignMetadata.Description, &campaignMetadata.Image); err != nil {
			return models.CampaignMetadata{}, err
		}
	}
	return campaignMetadata, nil
}

func GetDonationsByDonor(owner []byte) ([]dtos.DonationViewDTO, error) {
	pool, err := connect()
	if err != nil {
		return nil, err
	}
	defer pool.Close()

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

func GetAvailableRefundsByDonor(donor []byte) ([]dtos.RefundViewDTO, error) {
	pool, err := connect()
	if err != nil {
		return nil, err
	}
	defer pool.Close()

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
    WHERE c.deadline_ts < EXTRACT(EPOCH FROM now())
    AND c.target_amount > 0
    AND d.donor = $1
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
