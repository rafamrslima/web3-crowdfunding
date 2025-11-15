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
	"web3crowdfunding/internal/utils"

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

func GetContractAddress() (string, error) {
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	if contractAddress == "" {
		log.Printf("CONTRACT_ADDRESS environment variable not set")
		return "", fmt.Errorf("CONTRACT_ADDRESS not found")
	}
	return contractAddress, nil
}

func connectToEthereumNode() (*ethclient.Client, error) {
	ethClient, err := ethclient.Dial(ethClientAddress)
	if err != nil {
		log.Printf("Error connecting to Ethereum client: %v", err)
		return nil, err
	}
	return ethClient, nil
}

func loadPrivateKeyFromEnv() (*ecdsa.PrivateKey, error) {
	key := os.Getenv("PRIVATE_KEY")
	key = strings.TrimSpace(key)
	key = strings.TrimPrefix(key, "0x")
	key = strings.TrimPrefix(key, "0X")

	privateKey, err := crypto.HexToECDSA(key)
	return privateKey, err
}

func getContractABIFilePath() string {
	abiPath := os.Getenv("CROWDFUNDING_ABI_PATH")
	if abiPath == "" {
		return defaultABIPath
	}
	return abiPath
}

func parseContractABI() (abi.ABI, error) {
	abiPath := getContractABIFilePath()
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

func initializeCrowdfundingContract() (crowdfunding.Crowdfunding, error) {
	contractAddress, err := GetContractAddress()

	if err != nil {
		return crowdfunding.Crowdfunding{}, err
	}

	ethClient, err := connectToEthereumNode()

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

func BuildCampaignTransaction(campaign dtos.CampaignDto) (dtos.UnsignedTxResponse, error) {
	parsedABI, err := parseContractABI()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	weiTarget, err := utils.ParseEthToWei(campaign.Target)
	if err != nil {
		log.Printf("error: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}
	fmt.Printf("wei target: %v", weiTarget)

	deadline, _ := new(big.Int).SetString(campaign.Deadline, 10)
	data, err := parsedABI.Pack("createCampaign", campaign.Owner, campaign.Title, campaign.Description, weiTarget, deadline, campaign.Image)
	if err != nil {
		log.Printf("Error packing function data: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddress, err := GetContractAddress()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	ethClient, err := connectToEthereumNode()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
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
		return dtos.UnsignedTxResponse{}, err
	}

	unsigned := dtos.UnsignedTxResponse{
		To:    contractAddr.Hex(),
		Data:  fmt.Sprintf("0x%x", data),
		Value: "0x0", // no ETH being sent here
		Gas:   fmt.Sprintf("0x%x", gas),
	}

	return unsigned, nil
}

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

	weiTarget, err := utils.ParseEthToWei(campaign.Target)
	fmt.Printf("wei target: %v", weiTarget)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	deadline, _ := new(big.Int).SetString(campaign.Deadline, 10)
	transaction, err := contract.CreateCampaign(auth, campaign.Owner, campaign.Title, campaign.Description, weiTarget, deadline, campaign.Image)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return transaction, nil
}

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

func BuildDonationTransaction(campaignId int, value string) (dtos.UnsignedTxResponse, error) {
	parsedABI, err := parseContractABI()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	fmt.Printf("campaign id: %v", campaignId)

	campaignIdBigInt := big.NewInt(int64(campaignId))
	data, err := parsedABI.Pack("donateToCampaign", campaignIdBigInt)
	if err != nil {
		log.Printf("Error packing function data: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddress, err := GetContractAddress()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	ethClient, err := connectToEthereumNode()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddr := common.HexToAddress(contractAddress)
	defer ethClient.Close()

	valueInWei, err := utils.ParseEthToWei(value)
	fmt.Printf("wei parsed: %v /n", valueInWei)
	if err != nil {
		log.Printf("Error: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}

	callMsg := geth.CallMsg{
		To:    &contractAddr,
		Data:  data,
		Value: valueInWei,
	}

	gas, err := ethClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Printf("Error estimating gas: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}

	unsigned := dtos.UnsignedTxResponse{
		To:    contractAddr.Hex(),
		Data:  fmt.Sprintf("0x%x", data),
		Value: fmt.Sprintf("0x%x", valueInWei),
		Gas:   fmt.Sprintf("0x%x", gas),
	}

	fmt.Printf("object: %v", unsigned)

	return unsigned, nil
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

	valueWei, err := utils.ParseEthToWei(value)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	auth.Value = valueWei

	campaignIdBigInt := big.NewInt(int64(campaignId))
	transaction, err := contract.DonateToCampaign(auth, campaignIdBigInt)

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return transaction, nil
}
