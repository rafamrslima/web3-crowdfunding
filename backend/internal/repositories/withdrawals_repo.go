package repositories

import (
	"context"
	"log"
	"web3crowdfunding/internal/models"
)

func SaveWithdrawCompletion(withdraw models.WithdrawDbEntity) error {
	pool, err := GetDB()
	if err != nil {
		return err
	}

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
