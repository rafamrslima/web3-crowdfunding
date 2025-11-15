package dtos

type DonationDTO struct {
	CampaignId int    `json:"campaignId"`
	Value      string `json:"value"`
}
