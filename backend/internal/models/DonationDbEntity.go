package models

type DonationDbEntity struct {
	Id           int64
	CampaignId   int64
	Donor        string
	AmountWei    int64
	TxHash       string
	CreatedBlock uint64
	CreatedAt    uint64
}
