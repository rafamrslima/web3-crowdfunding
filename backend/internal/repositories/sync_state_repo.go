package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

func GetLastProcessedBlock(chainID int64) (uint64, error) {
	pool, err := GetDB()
	if err != nil {
		return 0, err
	}

	ctx := context.Background()
	var lastBlock int64

	err = pool.QueryRow(ctx,
		`SELECT last_processed_block FROM sync_state WHERE chain_id = $1`,
		chainID).Scan(&lastBlock)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// No record exists yet - this is first run
			log.Printf("No sync state found for chain %d, starting from block 0", chainID)
			return 0, nil
		}
		log.Printf("Error querying sync_state: %v", err)
		return 0, err
	}

	return uint64(lastBlock), nil
}

func UpdateLastProcessedBlock(chainID int64, blockNumber uint64) error {
	pool, err := GetDB()
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = pool.Exec(ctx,
		`INSERT INTO sync_state (chain_id, last_processed_block) 
		VALUES ($1, $2)
		ON CONFLICT (chain_id) 
		DO UPDATE SET last_processed_block = $2`,
		chainID, blockNumber)

	if err != nil {
		log.Printf("Error updating sync_state: %v", err)
		return err
	}

	log.Printf("Updated sync_state for chain %d to block %d", chainID, blockNumber)
	return nil
}
