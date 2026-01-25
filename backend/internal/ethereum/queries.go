package ethereum

import (
	"fmt"
	"log"
	crowdfunding "web3crowdfunding/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func FetchAllCampaigns() ([]crowdfunding.CrowdFundingCampaign, error) {
	contract, err := initializeCrowdfundingContract()
	if err != nil {
		return []crowdfunding.CrowdFundingCampaign{}, err
	}

	campaigns, err := contract.GetCampaigns(&bind.CallOpts{})
	if err != nil {
		log.Printf("error: %v", err)
		return []crowdfunding.CrowdFundingCampaign{}, err
	}

	return campaigns, nil
}

func FetchCampaignById(campaignId int) (crowdfunding.CrowdFundingCampaign, error) {
	contract, err := initializeCrowdfundingContract()
	if err != nil {
		return crowdfunding.CrowdFundingCampaign{}, err
	}

	campaigns, err := contract.GetCampaigns(&bind.CallOpts{})
	if err != nil {
		log.Printf("error: %v", err)
		return crowdfunding.CrowdFundingCampaign{}, err
	}

	if campaignId+1 > len(campaigns) {
		return crowdfunding.CrowdFundingCampaign{}, fmt.Errorf("campaign ID %d not found", campaignId)
	}
	return campaigns[campaignId], nil
}
