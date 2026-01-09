package dtos

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type DonationDTO struct {
	CampaignId int    `json:"campaignId"`
	Value      string `json:"value"`
}

type DonationViewDTO struct {
	Donor       common.Address `json:"donor"`
	CampaignId  string         `json:"campaignId"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	Image       string         `json:"image"`
	Amount      int64          `json:"amount"`
}
