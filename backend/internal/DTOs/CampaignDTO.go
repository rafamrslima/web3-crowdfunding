package dtos

import (
	"github.com/ethereum/go-ethereum/common"
)

type CampaignDto struct {
	Owner           common.Address `json:"owner"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	Target          string         `json:"target"`
	Deadline        string         `json:"deadline"`
	Image           string         `json:"image"`
	AmountCollected *uint64        `json:"amountCollected"`
}
