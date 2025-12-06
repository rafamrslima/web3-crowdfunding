package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/models"

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

func SaveCampaignCreated(campaign models.CampaignDbEntity) error {
	pool, err := connect()
	if err != nil {
		return err
	}
	defer pool.Close()

	ctx := context.Background()

	_, err = pool.Exec(ctx,
		`INSERT INTO campaigns (campaign_id, owner, target_amount, deadline_ts, tx_hash, block_number, block_time) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		campaign.Id, campaign.Owner, campaign.Target, campaign.Deadline, campaign.TxHash, campaign.BlockNumber, campaign.BlockTime)

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
	rows, err := pool.Query(ctx, `SELECT owner, target_amount, deadline_ts FROM campaigns WHERE owner = $1`, owner)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []dtos.CampaignDto

	for rows.Next() {
		var res dtos.CampaignDto
		if err := rows.Scan(&res.Owner, &res.Target, &res.Deadline); err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}
