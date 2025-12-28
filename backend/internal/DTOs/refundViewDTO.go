package dtos

type RefundViewDTO struct {
	DonationView            DonationViewDTO `json:"donationView"`
	CampaignTarget          string          `json:"target"`
	CampaignDeadline        string          `json:"deadline"`
	CampaignAmountCollected string          `json:"amountCollected"`
}
