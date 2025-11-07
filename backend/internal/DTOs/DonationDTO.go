package dtos

import "math/big"

type DonationDTO struct {
	CampaignId big.Int `json:"campaignId"`
	Value      int64   `json:"value"`
}
