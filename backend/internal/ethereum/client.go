package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"strings"
	crowdfunding "web3crowdfunding/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
