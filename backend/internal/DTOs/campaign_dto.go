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
	CategoryId      *int32         `json:"categoryId,omitempty"`
	AmountCollected *uint64        `json:"amountCollected"`
}
