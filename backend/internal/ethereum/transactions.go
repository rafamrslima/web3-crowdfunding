package ethereum

import (
	"fmt"
	"log"
	"math/big"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/repositories"
	"web3crowdfunding/internal/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func BuildCampaignTransaction(campaign dtos.CampaignDto) (dtos.UnsignedTxResponse, error) {
	parsedABI, err := parseContractABI()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	target, err := utils.ParseUSDC(campaign.Target)
	if err != nil {
		log.Printf("error: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}

	deadline, _ := new(big.Int).SetString(campaign.Deadline, 10)
	creationIdStr := uuid.NewString()
	creationIdHash := crypto.Keccak256Hash([]byte(creationIdStr))

	data, err := parsedABI.Pack("createCampaign", target, deadline, creationIdHash)
	if err != nil {
		log.Printf("Error packing function data: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddress, err := GetContractAddress()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddr := common.HexToAddress(contractAddress)

	err = repositories.SaveCampaignDraft(creationIdHash.Hex(),
		campaign.Owner, campaign.Title, campaign.Description, campaign.Image, campaign.CategoryId)

	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	creationId := creationIdHash.Hex()
	unsigned := dtos.UnsignedTxResponse{
		To:         contractAddr.Hex(),
		Data:       fmt.Sprintf("0x%x", data),
		Value:      "0x0",
		Gas:        fmt.Sprintf("0x%x", defaultGasEstimation),
		CreationId: &creationId,
	}

	return unsigned, nil
}

func BuildDonationTransaction(campaignId int, value string) (dtos.UnsignedTxResponse, error) {
	parsedABI, err := parseContractABI()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	campaignIdBigInt := big.NewInt(int64(campaignId))
	valueParsed, _ := utils.ParseUSDC(value)
	data, err := parsedABI.Pack("donateToCampaign", campaignIdBigInt, valueParsed)
	if err != nil {
		log.Printf("Error packing function data: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddress, err := GetContractAddress()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddr := common.HexToAddress(contractAddress)

	unsigned := dtos.UnsignedTxResponse{
		To:    contractAddr.Hex(),
		Data:  fmt.Sprintf("0x%x", data),
		Value: "0x0",
		Gas:   fmt.Sprintf("0x%x", defaultGasEstimation),
	}

	return unsigned, nil
}

func BuildWithdrawTransaction(campaignId int) (dtos.UnsignedTxResponse, error) {
	parsedABI, err := parseContractABI()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	campaignIdBigInt := big.NewInt(int64(campaignId))
	data, err := parsedABI.Pack("withdraw", campaignIdBigInt)
	if err != nil {
		log.Printf("Error packing function data: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddress, err := GetContractAddress()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddr := common.HexToAddress(contractAddress)

	unsigned := dtos.UnsignedTxResponse{
		To:    contractAddr.Hex(),
		Data:  fmt.Sprintf("0x%x", data),
		Value: "0x0",
		Gas:   fmt.Sprintf("0x%x", defaultGasEstimation),
	}

	return unsigned, nil
}

func BuildRefundTransaction(campaignId int) (dtos.UnsignedTxResponse, error) {
	parsedABI, err := parseContractABI()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	campaignIdBigInt := big.NewInt(int64(campaignId))
	data, err := parsedABI.Pack("refundDonor", campaignIdBigInt)
	if err != nil {
		log.Printf("Error packing function data: %v", err)
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddress, err := GetContractAddress()
	if err != nil {
		return dtos.UnsignedTxResponse{}, err
	}

	contractAddr := common.HexToAddress(contractAddress)

	unsigned := dtos.UnsignedTxResponse{
		To:    contractAddr.Hex(),
		Data:  fmt.Sprintf("0x%x", data),
		Value: "0x0",
		Gas:   fmt.Sprintf("0x%x", defaultGasEstimation),
	}

	return unsigned, nil
}
