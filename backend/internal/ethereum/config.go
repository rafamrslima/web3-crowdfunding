package ethereum

import "os"

const (
	defaultABIPath       = "contracts/crowdfunding.abi"
	defaultGasEstimation = 250000
)

func getEthClientAddress() string {
	if addr := os.Getenv("ETH_RPC_URL"); addr != "" {
		return addr
	}
	return "http://127.0.0.1:8545" // fallback to local
}
