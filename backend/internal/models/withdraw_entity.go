package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type WithdrawDbEntity struct {
	CampaignId  int64
	Owner       common.Address
	Amount      int64
	TxHash      common.Hash
	BlockNumber uint64
	BlockTime   time.Time
	CreatedAt   time.Time
}
