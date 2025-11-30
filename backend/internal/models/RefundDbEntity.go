package models

import "time"

type RefundDbEntity struct {
	CampaignId       int64
	Donor            string
	TotalContributed int64
	TxHash           string
	BlockNumber      uint64
	BlockTime        time.Time
	CreatedAt        time.Time
}
