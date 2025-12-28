package dtos

type RefundViewDTO struct {
	Donation                DonationViewDTO `json:"donation"`
	CampaignTarget          string          `json:"target"`
	CampaignDeadline        string          `json:"deadline"`
	CampaignAmountCollected string          `json:"amountCollected"`
}
