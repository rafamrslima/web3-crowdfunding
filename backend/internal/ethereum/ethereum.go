package ethereum

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"strings"
	crowdfunding "web3crowdfunding/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func loadContract() (crowdfunding.Crowdfunding, error) {
	contractAddress := os.Getenv("CONTRACT_ADDRESS")

	client, err := ethclient.Dial("http://127.0.0.1:8545")

	if err != nil {
		log.Printf("error: %v", err)
		return crowdfunding.Crowdfunding{}, err
	}

	contractAddr := common.HexToAddress(contractAddress)
	contract, err := crowdfunding.NewCrowdfunding(contractAddr, client)

	if err != nil {
		log.Printf("error: %v", err)
		return crowdfunding.Crowdfunding{}, err
	}

	return *contract, nil
}

func CreateCampaign(owner common.Address, title string, description string, target *big.Int, deadline *big.Int, image string) (*types.Transaction, error) {
	contract, err := loadContract()

	if err != nil {
		return nil, err
	}

	privateKey, err := getPrivateKey()
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	chainID := big.NewInt(31337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	if err != nil {
		return nil, err
	}

	transaction, err := contract.CreateCampaign(auth, owner, title, description, target, deadline, image)

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return transaction, nil
}

func GetCampaigns() ([]crowdfunding.CrowdFundingCampaign, error) {
	contract, err := loadContract()

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

func DonateToCampaign(campaignId big.Int, value int64) (*types.Transaction, error) {
	contract, err := loadContract()

	if err != nil {
		return nil, err
	}

	privateKey, err := getPrivateKey()
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	chainID := big.NewInt(31337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	if err != nil {
		return nil, err
	}

	auth.Value = big.NewInt(value)
	transaction, err := contract.DonateToCampaign(auth, &campaignId)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func getPrivateKey() (*ecdsa.PrivateKey, error) {
	key := os.Getenv("PRIVATE_KEY")
	key = strings.TrimSpace(key)
	key = strings.TrimPrefix(key, "0x")
	key = strings.TrimPrefix(key, "0X")

	privateKey, err := crypto.HexToECDSA(key)
	return privateKey, err
}
