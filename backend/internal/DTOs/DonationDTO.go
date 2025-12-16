package dtos

type DonationDTO struct {
	CampaignId int    `json:"campaignId"`
	Value      string `json:"value"`
}

type DonationViewDTO struct {
	Donor       string `json:"donor"`
	CampaignId  string `json:"campaignId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	Image       string `json:"image"`
	Amount      int    `json:"amount"`
}
