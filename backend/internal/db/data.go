package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
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
		`INSERT INTO campaigns (campaignId, owner, target_wei, deadline_ts, created_tx, created_block, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		campaign.Id, campaign.Owner, campaign.Target, campaign.Deadline, campaign.CampaignTx, campaign.Block, time.Now())

	if err != nil {
		return err
	}

	log.Println("Row inserted successfully.")
	return nil
}
