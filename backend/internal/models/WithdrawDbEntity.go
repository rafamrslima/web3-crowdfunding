package models

import "time"

type WithdrawDbEntity struct {
	CampaignId  int64
	Owner       string
	Amount      int64
	TxHash      string
	BlockNumber uint64
	BlockTime   time.Time
	CreatedAt   time.Time
}
