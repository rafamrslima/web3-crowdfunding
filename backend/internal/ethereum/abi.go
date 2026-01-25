package ethereum

import (
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

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
