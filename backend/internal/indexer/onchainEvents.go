package indexer

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"web3crowdfunding/internal/db"
	internalEthereum "web3crowdfunding/internal/ethereum"
	"web3crowdfunding/internal/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ethClientAddress = "//127.0.0.1:8545"
const defaultABIPath = "contracts/crowdfunding.abi"

func startHttpConnection(ctx context.Context) *ethclient.Client {
	httpURL := "http:" + ethClientAddress

	httpClient, err := ethclient.DialContext(ctx, httpURL)
	if err != nil {
		log.Fatal("http dial:", err)
	}
	return httpClient
}

func startWebSocketConnection(ctx context.Context) *ethclient.Client {
	wsURL := "ws:" + ethClientAddress

	wsClient, err := ethclient.DialContext(ctx, wsURL)
	if err != nil {
		log.Fatal("websocket dial:", err)
	}
	return wsClient
}

func StartEventListener() {
	fmt.Println("starting listener...")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	contractAddress, err := internalEthereum.GetContractAddress()
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress(contractAddress)

	httpClient := startHttpConnection(ctx)
	defer httpClient.Close()

	wsClient := startWebSocketConnection(ctx)
	defer wsClient.Close()

	abiBytes, err := os.ReadFile(defaultABIPath)
	if err != nil {
		log.Fatal("read abi:", err)
	}

	parsedABI, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		log.Fatal("parse abi:", err)
	}

	events, ok := parsedABI.Events["CampaignCreated"]
	if !ok {
		log.Fatal("event CampaignCreated not found in ABI")
	}
	topic0 := events.ID

	ch := make(chan types.Log)
	sub, err := wsClient.SubscribeFilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
		Topics:    [][]common.Hash{{topic0}},
	}, ch)
	if err != nil {
		log.Fatal("subscribe:", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("shutting down listener")
			return

		case err := <-sub.Err():
			log.Println("subscription error:", err)
			return

		case lg := <-ch:
			printCampaignCreated(parsedABI, lg)
		}
	}
}

func printCampaignCreated(parsedABI abi.ABI, lg types.Log) {
	id := new(big.Int).SetBytes(lg.Topics[1].Bytes())
	owner := common.BytesToAddress(lg.Topics[2].Bytes())

	var out struct {
		Title     string
		TargetWei *big.Int
		Deadline  *big.Int
	}

	if err := parsedABI.UnpackIntoInterface(&out, "CampaignCreated", lg.Data); err != nil {
		log.Println("unpack:", err)
		return
	}

	fmt.Printf("CampaignCreated id=%s owner=%s title=%s targetWei=%s deadline=%d block=%d\n",
		id.String(),
		owner.Hex(),
		out.Title,
		out.TargetWei.String(),
		out.Deadline.Uint64(),
		lg.BlockNumber,
	)

	campaignDbObj := models.CampaignDbEntity{
		Id:       id.Int64(),
		Owner:    owner.Hex(),
		Title:    out.Title,
		Target:   out.TargetWei.String(),
		Deadline: uint64(out.Deadline.Int64()),
		Block:    lg.BlockNumber,
	}

	err := db.SaveCreatedCampaign(campaignDbObj)
	if err != nil {
		log.Println(err)
		return
	}
}
