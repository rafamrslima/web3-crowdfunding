package validation

import (
	"errors"
	"math/big"
	"strings"
	dtos "web3crowdfunding/internal/DTOs"

	"github.com/ethereum/go-ethereum/common"
)

func ValidateCampaign(campaign dtos.CampaignDto) error {
	if campaign.Owner == (common.Address{}) {
		return errors.New("owner address is required and cannot be empty")
	}

	if strings.TrimSpace(campaign.Title) == "" {
		return errors.New("title is required and cannot be empty")
	}
	if len(campaign.Title) > 100 {
		return errors.New("title cannot exceed 100 characters")
	}

	if strings.TrimSpace(campaign.Description) == "" {
		return errors.New("description is required and cannot be empty")
	}
	if len(campaign.Description) > 1000 {
		return errors.New("description cannot exceed 1000 characters")
	}

	if campaign.Target.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("target must be greater than 0")
	}

	if campaign.Deadline.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("deadline must be a positive timestamp")
	}

	if campaign.Image != "" && strings.TrimSpace(campaign.Image) == "" {
		return errors.New("image URL cannot be empty if provided")
	}

	return nil
}
