package db

import (
	"context"
	"fmt"
	"log"
	"os"
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

func SaveTempCampaignMetadata(address common.Address, nonce uint64, title string, description string, image string) error {
	pool, err := connect()
	if err != nil {
		return err
	}
	defer pool.Close()

	ctx := context.Background()

	_, err = pool.Exec(ctx,
		`INSERT INTO tempCampaignMetadata (owner, nonce, title, description, image) 
		VALUES ($1, $2, $3, $4, $5)`,
		address, nonce, title, description, image)

	if err != nil {
		return err
	}

	log.Println("Row inserted successfully.")
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
		return err
	}

	log.Println("Row inserted successfully.")
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
		return err
	}

	log.Println("Row inserted successfully.")
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
		return err
	}

	log.Println("Row inserted successfully.")
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
		return err
	}

	log.Println("Row inserted successfully.")
	return nil
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

func GetTempCampaignMetadata(owner common.Address, nonce uint64) (models.CampaignMetadata, error) {
	pool, err := connect()
	if err != nil {
		return models.CampaignMetadata{}, err
	}
	defer pool.Close()

	ctx := context.Background()
	rows, err := pool.Query(ctx, `SELECT title, description, image FROM tempCampaignMetadata WHERE owner = $1 and nonce = $2 LIMIT 1`, owner, nonce)

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
