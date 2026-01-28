package ethereum

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetCampaignsTotal() (uint64, error) {
	contract, err := initializeCrowdfundingContract()
	if err != nil {
		return 0, err
	}

	total, err := contract.GetCampaignsTotal(&bind.CallOpts{})
	if err != nil {
		log.Printf("error: %v", err)
		return 0, err
	}

	return total.Uint64(), nil
}
