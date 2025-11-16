package models

type CampaignDbEntity struct {
	Id         int64
	Owner      string
	Title      string
	Target     string
	Deadline   uint64
	CampaignTx string
	Block      uint64
}
