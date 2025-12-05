package models

import "time"

type DonationDbEntity struct {
	Id          int64
	CampaignId  int64
	Donor       string
	Amount      int64
	TxHash      string
	BlockNumber uint64
	BlockTime   time.Time
	CreatedAt   time.Time
}
