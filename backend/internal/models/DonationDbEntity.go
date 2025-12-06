package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type DonationDbEntity struct {
	Id          int64
	CampaignId  int64
	Donor       common.Address
	Amount      int64
	TxHash      common.Hash
	BlockNumber uint64
	BlockTime   time.Time
	CreatedAt   time.Time
}
