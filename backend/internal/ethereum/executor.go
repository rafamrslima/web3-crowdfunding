package ethereum

import (
	"log"
	"math/big"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func ExecuteCampaignCreation(campaign dtos.CampaignDto) (*types.Transaction, error) {
	contract, err := initializeCrowdfundingContract()
	if err != nil {
		return nil, err
	}

	privateKey, err := loadPrivateKeyFromEnv()
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	chainID := big.NewInt(31337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	target, err := utils.ParseUSDC(campaign.Target)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	deadline, _ := new(big.Int).SetString(campaign.Deadline, 10)
	creationIdStr := uuid.NewString()
	creationId := crypto.Keccak256Hash([]byte(creationIdStr))
	transaction, err := contract.CreateCampaign(auth, target, deadline, creationId)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return transaction, nil
}

func ExecuteDonationToCompaign(campaignId int, value string) (*types.Transaction, error) {
	contract, err := initializeCrowdfundingContract()
	if err != nil {
		return nil, err
	}

	privateKey, err := loadPrivateKeyFromEnv()
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	chainID := big.NewInt(31337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	valueParsed, err := utils.ParseUSDC(value)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	campaignIdBigInt := big.NewInt(int64(campaignId))
	transaction, err := contract.DonateToCampaign(auth, campaignIdBigInt, valueParsed)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return transaction, nil
}
