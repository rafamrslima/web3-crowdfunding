package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type CampaignDbEntity struct {
	Id          int64
	Owner       common.Address
	Title       string
	Target      int64
	Deadline    uint64
	TxHash      common.Hash
	BlockNumber uint64
	BlockTime   time.Time
	CreatedAt   time.Time
}
