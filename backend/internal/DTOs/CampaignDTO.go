package dtos

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type CampaignDto struct {
	Owner       common.Address `json:"owner"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Target      big.Int        `json:"target"`
	Deadline    big.Int        `json:"deadline"`
	Image       string         `json:"image"`
}
