package indexer

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"sync"
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

	contractAddress, err := internalEthereum.GetContractAddress()
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress(contractAddress)
	wsClient := startWebSocketConnection(ctx)

	abiBytes, err := os.ReadFile(defaultABIPath)
	if err != nil {
		log.Fatal("read abi:", err)
	}

	parsedABI, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		log.Fatal("parse abi:", err)
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		listenToCampaignCreation(contractAddr, parsedABI, ctx, wsClient)
	}()

	go func() {
		defer wg.Done()
		listenToDonationCreation(contractAddr, parsedABI, ctx, wsClient)
	}()

	go func() {
		defer wg.Done()
		listenToRefundIssued(contractAddr, parsedABI, ctx, wsClient)
	}()

	go func() {
		<-ctx.Done()
		fmt.Println("Shutdown signal received, closing WebSocket...")
		wsClient.Close()
	}()

	wg.Wait()
	stop()
}

func listenToCampaignCreation(contractAddr common.Address, parsedABI abi.ABI, ctx context.Context, wsClient *ethclient.Client) {
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
			SaveCampaignCreated(parsedABI, lg)
		}
	}
}

func listenToDonationCreation(contractAddr common.Address, parsedABI abi.ABI, ctx context.Context, wsClient *ethclient.Client) {
	events, ok := parsedABI.Events["DonationReceived"]
	if !ok {
		log.Fatal("event DonationReceived not found in ABI")
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
			saveDonationReceived(parsedABI, lg)
		}
	}
}

func listenToRefundIssued(contractAddr common.Address, parsedABI abi.ABI, ctx context.Context, wsClient *ethclient.Client) {
	events, ok := parsedABI.Events["RefundIssued"]
	if !ok {
		log.Fatal("event RefundIssued not found in ABI")
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
			saveRefundIssued(parsedABI, lg)
		}
	}
}

func SaveCampaignCreated(parsedABI abi.ABI, lg types.Log) {
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

	campaignDbObj := models.CampaignDbEntity{
		Id:         id.Int64(),
		Owner:      owner.Hex(),
		Title:      out.Title,
		Target:     out.TargetWei.Int64(),
		Deadline:   uint64(out.Deadline.Int64()),
		CampaignTx: lg.TxHash.Hex(),
		Block:      lg.BlockNumber,
	}

	err := db.SaveCampaignCreated(campaignDbObj)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("CampaignCreated id=%s owner=%s title=%s targetWei=%s txHash=%s deadline=%d block=%d\n",
		id.String(),
		owner.Hex(),
		out.Title,
		out.TargetWei.String(),
		lg.TxHash,
		out.Deadline.Uint64(),
		lg.BlockNumber,
	)
}

func saveDonationReceived(parsedABI abi.ABI, lg types.Log) {
	campaignId := new(big.Int).SetBytes(lg.Topics[1].Bytes())
	donor := common.BytesToAddress(lg.Topics[2].Bytes())

	var out struct {
		AmountWei *big.Int
	}

	if err := parsedABI.UnpackIntoInterface(&out, "DonationReceived", lg.Data); err != nil {
		log.Println("unpack:", err)
		return
	}

	donationDbObj := models.DonationDbEntity{
		CampaignId:   campaignId.Int64(),
		Donor:        donor.Hex(),
		AmountWei:    out.AmountWei.Int64(),
		TxHash:       lg.TxHash.Hex(),
		CreatedBlock: lg.BlockNumber,
	}

	err := db.SaveDonationReceived(donationDbObj)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("DonationReceived campaignId=%s donor=%s amountWei=%s txHash=%s block=%d\n",
		campaignId.String(),
		donor.Hex(),
		out.AmountWei.String(),
		lg.TxHash,
		lg.BlockNumber,
	)
}

func saveRefundIssued(parsedABI abi.ABI, lg types.Log) {
}
