package main

import (
	"fmt"
	"log"
	"os"
	crowdfunding "web3crowdfunding/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello web3")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env.")
	}

	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	fmt.Println(contractAddress)

	client, err := ethclient.Dial("http://127.0.0.1:8545")

	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress(contractAddress)
	contract, err := crowdfunding.NewCrowdfunding(contractAddr, client)

	if err != nil {
		log.Fatal(err)
	}

	campaigns, err := contract.GetCampaigns(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Campaigns:", campaigns)
}
