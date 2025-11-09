package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	crowdfunding "web3crowdfunding/contracts"
	dtos "web3crowdfunding/internal/DTOs"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const defaultABIPath = "contracts/crowdfunding.abi"
const ethClientAddress = "http://127.0.0.1:8545"

func loadContract() (string, error) {
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	if contractAddress == "" {
		log.Printf("CONTRACT_ADDRESS environment variable not set")
		return "", fmt.Errorf("CONTRACT_ADDRESS not found")
	}
	return contractAddress, nil
}

func loadEthClient() (*ethclient.Client, error) {
	ethClient, err := ethclient.Dial(ethClientAddress)
	if err != nil {
		log.Printf("Error connecting to Ethereum client: %v", err)
		return nil, err
	}
	return ethClient, nil
}

func loadCrowdfundingClient() (crowdfunding.Crowdfunding, error) {
	contractAddress, err := loadContract()

	if err != nil {
		return crowdfunding.Crowdfunding{}, err
	}

	ethClient, err := loadEthClient()

	if err != nil {
		return crowdfunding.Crowdfunding{}, err
	}

	contractAddr := common.HexToAddress(contractAddress)
	contract, err := crowdfunding.NewCrowdfunding(contractAddr, ethClient)

	if err != nil {
		log.Printf("error: %v", err)
		return crowdfunding.Crowdfunding{}, err
	}

	return *contract, nil
}

func CreateUnsignedCampaign(campaign dtos.CampaignDto) (dtos.UnsignedTx, error) {
	parsedABI, err := getParsedABI()
	if err != nil {
		return dtos.UnsignedTx{}, err
	}

	data, err := parsedABI.Pack("createCampaign", campaign.Owner, campaign.Title, campaign.Description, &campaign.Target, &campaign.Deadline, campaign.Image)
	if err != nil {
		log.Printf("Error packing function data: %v", err)
		return dtos.UnsignedTx{}, err
	}

	contractAddress, err := loadContract()
	if err != nil {
		return dtos.UnsignedTx{}, err
	}

	ethClient, err := loadEthClient()
	if err != nil {
		return dtos.UnsignedTx{}, err
	}

	contractAddr := common.HexToAddress(contractAddress)
	defer ethClient.Close()

	callMsg := geth.CallMsg{
		To:   &contractAddr,
		Data: data,
	}

	gas, err := ethClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Printf("Error estimating gas: %v", err)
		return dtos.UnsignedTx{}, err
	}

	unsigned := dtos.UnsignedTx{
		To:    contractAddr.Hex(),
		Data:  fmt.Sprintf("0x%x", data),
		Value: "0x0", // no ETH being sent here
		Gas:   fmt.Sprintf("0x%x", gas),
	}

	return unsigned, nil
}

func CreateCampaign(owner common.Address, title string, description string, target *big.Int, deadline *big.Int, image string) (*types.Transaction, error) {
	contract, err := loadCrowdfundingClient()

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
	contract, err := loadCrowdfundingClient()

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
	contract, err := loadCrowdfundingClient()

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

func getABIPath() string {
	abiPath := os.Getenv("CROWDFUNDING_ABI_PATH")
	if abiPath == "" {
		return defaultABIPath
	}
	return abiPath
}

func getParsedABI() (abi.ABI, error) {
	abiPath := getABIPath()
	abiBytes, err := os.ReadFile(abiPath)
	if err != nil {
		log.Printf("Error reading ABI file from %s: %v", abiPath, err)
		return abi.ABI{}, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		log.Printf("Error parsing ABI: %v", err)
		return abi.ABI{}, err
	}
	return parsedABI, nil
}
