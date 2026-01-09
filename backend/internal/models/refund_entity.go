package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type RefundDbEntity struct {
	CampaignId       int64
	Donor            common.Address
	TotalContributed int64
	TxHash           common.Hash
	BlockNumber      uint64
	BlockTime        time.Time
	CreatedAt        time.Time
}
