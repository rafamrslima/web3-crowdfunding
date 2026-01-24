package repositories

import (
	"context"
	"log"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/models"

	"github.com/ethereum/go-ethereum/common"
)

func SaveCampaignDraft(creationId string, address common.Address, title string, description string, image string, categoryId *int32) error {
	pool, err := GetDB()
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, err = pool.Exec(ctx,
		`INSERT INTO campaign_drafts (creation_id, owner, title, description, image, category_id) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		creationId, address, title, description, image, categoryId)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Row inserted successfully into campaign_drafts.")
	return nil
}

func SaveCampaignCreated(campaign models.CampaignDbEntity) error {
	pool, err := GetDB()
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, err = pool.Exec(ctx,
		`INSERT INTO campaigns 
		(campaign_id, owner, title, description, target_amount, deadline_ts, image, category_id, tx_hash, block_number, block_time) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		campaign.Id,
		campaign.Owner,
		campaign.Title,
		campaign.Description,
		campaign.Target,
		campaign.Deadline,
		campaign.Image,
		campaign.CategoryId,
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

func FetchAllCampaigns() ([]dtos.CampaignDto, error) {
	pool, err := GetDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	rows, err := pool.Query(ctx,
		`SELECT c.owner, c.title, c.description, c.target_amount, c.deadline_ts, c.image, c.category_id, sum(d.amount) as amount_collected
		FROM campaigns c left join donations d on c.campaign_id = d.campaign_id 
	    GROUP BY c.campaign_id, c.owner, c.title, c.description, c.target_amount, c.deadline_ts, c.image, c.category_id 
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
		if err := rows.Scan(&res.Owner, &res.Title, &res.Description, &res.Target, &res.Deadline, &res.Image, &res.CategoryId, &res.AmountCollected); err != nil {
			log.Println(err)
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}

func GetCampaignsByOwner(owner []byte) ([]dtos.CampaignDto, error) {
	pool, err := GetDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	rows, err := pool.Query(ctx,
		`SELECT c.owner, c.title, c.description, c.target_amount, c.deadline_ts, c.image, c.category_id, sum(d.amount) as amount_collected
		FROM campaigns c left join donations d on c.campaign_id = d.campaign_id 
	    WHERE c.owner = $1 GROUP BY c.owner, c.title, c.description, c.target_amount, c.deadline_ts, c.image, c.category_id
		LIMIT 100`, owner)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var results []dtos.CampaignDto

	for rows.Next() {
		var res dtos.CampaignDto
		if err := rows.Scan(&res.Owner, &res.Title, &res.Description, &res.Target, &res.Deadline, &res.Image, &res.CategoryId, &res.AmountCollected); err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}

func GetCampaignMetadataFromDraft(owner common.Address, creationId string) (models.CampaignMetadata, error) {
	pool, err := GetDB()
	if err != nil {
		return models.CampaignMetadata{}, err
	}

	ctx := context.Background()
	rows, err := pool.Query(ctx,
		`SELECT title, description, image, category_id FROM campaign_drafts WHERE owner = $1 AND creation_id = $2 LIMIT 1`, owner, creationId)

	if err != nil {
		return models.CampaignMetadata{}, err
	}
	defer rows.Close()

	var campaignMetadata models.CampaignMetadata
	for rows.Next() {
		if err := rows.Scan(&campaignMetadata.Title, &campaignMetadata.Description, &campaignMetadata.Image, &campaignMetadata.CategoryId); err != nil {
			return models.CampaignMetadata{}, err
		}
	}
	return campaignMetadata, nil
}
