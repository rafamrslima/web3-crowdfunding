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
	"time"

	"web3crowdfunding/internal/database"
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
const campaignCreationEvent = "CampaignCreated"
const donationReceivedEvent = "DonationReceived"
const fundsWithdrawnEvent = "FundsWithdrawn"
const donationRefundedEvent = "DonationRefunded"

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
	wg.Add(4)

	go func() {
		defer wg.Done()
		listenToEventCreation(contractAddr, parsedABI, ctx, wsClient, campaignCreationEvent)
	}()

	go func() {
		defer wg.Done()
		listenToEventCreation(contractAddr, parsedABI, ctx, wsClient, donationReceivedEvent)
	}()

	go func() {
		defer wg.Done()
		listenToEventCreation(contractAddr, parsedABI, ctx, wsClient, fundsWithdrawnEvent)
	}()

	go func() {
		defer wg.Done()
		listenToEventCreation(contractAddr, parsedABI, ctx, wsClient, donationRefundedEvent)
	}()

	go func() {
		<-ctx.Done()
		fmt.Println("Shutdown signal received, closing WebSocket...")
		wsClient.Close()
	}()

	wg.Wait()
	stop()
}

func listenToEventCreation(contractAddr common.Address, parsedABI abi.ABI, ctx context.Context, wsClient *ethclient.Client, eventName string) {
	events, ok := parsedABI.Events[eventName]
	if !ok {
		log.Fatalf("event %v not found in ABI", eventName)
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
			switch eventName {
			case campaignCreationEvent:
				SaveCampaignCreated(parsedABI, lg)
			case donationReceivedEvent:
				saveDonationReceived(parsedABI, lg)
			case fundsWithdrawnEvent:
				saveWithdrawCompletion(parsedABI, lg)
			case donationRefundedEvent:
				saveDonationRefund(parsedABI, lg)
			default:
				log.Println("event not found.")
			}
		}
	}
}

func SaveCampaignCreated(parsedABI abi.ABI, lg types.Log) {
	id := new(big.Int).SetBytes(lg.Topics[1].Bytes())
	owner := common.BytesToAddress(lg.Topics[2].Bytes())
	creationId := lg.Topics[3]

	var out struct {
		Target   *big.Int
		Deadline *big.Int
	}

	if err := parsedABI.UnpackIntoInterface(&out, campaignCreationEvent, lg.Data); err != nil {
		log.Println("unpack:", err)
		return
	}

	campaignMetadata, err := database.GetCampaignMetadataFromDraft(owner, creationId.Hex())
	if err != nil {
		log.Println(err)
		return
	}

	campaignDbObj := models.CampaignDbEntity{
		Id:          id.Int64(),
		Owner:       owner,
		Title:       campaignMetadata.Title,
		Description: campaignMetadata.Description,
		Target:      out.Target.Int64(),
		Deadline:    uint64(out.Deadline.Int64()),
		Image:       campaignMetadata.Image,
		TxHash:      lg.TxHash,
		BlockNumber: lg.BlockNumber,
		BlockTime:   time.Unix(int64(lg.BlockTimestamp), 0),
	}

	err = database.SaveCampaignCreated(campaignDbObj)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("CampaignCreated id=%s owner=%s creationId=%s target=%s txHash=%s deadline=%d block=%d\n",
		id.String(),
		owner.Hex(),
		creationId,
		out.Target.String(),
		lg.TxHash,
		out.Deadline.Uint64(),
		lg.BlockNumber,
	)
}

func saveDonationReceived(parsedABI abi.ABI, lg types.Log) {
	campaignId := new(big.Int).SetBytes(lg.Topics[1].Bytes())
	donor := common.BytesToAddress(lg.Topics[2].Bytes())

	var out struct {
		Amount *big.Int
	}

	if err := parsedABI.UnpackIntoInterface(&out, donationReceivedEvent, lg.Data); err != nil {
		log.Println("unpack:", err)
		return
	}

	donationDbObj := models.DonationDbEntity{
		CampaignId:  campaignId.Int64(),
		Donor:       donor,
		Amount:      out.Amount.Int64(),
		TxHash:      lg.TxHash,
		BlockNumber: lg.BlockNumber,
		BlockTime:   time.Unix(int64(lg.BlockTimestamp), 0),
	}

	err := database.SaveDonationReceived(donationDbObj)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("DonationReceived campaignId=%s donor=%s amount=%s txHash=%s block=%d\n",
		campaignId.String(),
		donor.Hex(),
		out.Amount.String(),
		lg.TxHash,
		lg.BlockNumber,
	)
}

func saveWithdrawCompletion(parsedABI abi.ABI, lg types.Log) {
	campaignId := new(big.Int).SetBytes(lg.Topics[1].Bytes())
	owner := common.BytesToAddress(lg.Topics[2].Bytes())

	var out struct {
		Amount *big.Int
	}

	if err := parsedABI.UnpackIntoInterface(&out, fundsWithdrawnEvent, lg.Data); err != nil {
		log.Println("unpack:", err)
		return
	}

	withdrawDbObj := models.WithdrawDbEntity{
		CampaignId:  campaignId.Int64(),
		Owner:       owner,
		Amount:      out.Amount.Int64(),
		TxHash:      lg.TxHash,
		BlockNumber: lg.BlockNumber,
		BlockTime:   time.Unix(int64(lg.BlockTimestamp), 0),
	}

	err := database.SaveWithdrawCompletion(withdrawDbObj)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("FundsWithdrawn campaignId=%s owner=%s amount=%s txHash=%s block=%d\n",
		campaignId.String(),
		owner.Hex(),
		out.Amount.String(),
		lg.TxHash,
		lg.BlockNumber,
	)
}

func saveDonationRefund(parsedABI abi.ABI, lg types.Log) {
	campaignId := new(big.Int).SetBytes(lg.Topics[1].Bytes())
	donor := common.BytesToAddress(lg.Topics[2].Bytes())

	var out struct {
		TotalContributed *big.Int
	}

	if err := parsedABI.UnpackIntoInterface(&out, donationRefundedEvent, lg.Data); err != nil {
		log.Println("unpack:", err)
		return
	}

	refundDbObj := models.RefundDbEntity{
		CampaignId:       campaignId.Int64(),
		Donor:            donor,
		TotalContributed: out.TotalContributed.Int64(),
		TxHash:           lg.TxHash,
		BlockNumber:      lg.BlockNumber,
		BlockTime:        time.Unix(int64(lg.BlockTimestamp), 0),
	}

	err := database.SaveRefundIssued(refundDbObj)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Refund Issued campaignId=%s donor=%s totalContributed=%s txHash=%s block=%d\n",
		campaignId.String(),
		donor.Hex(),
		out.TotalContributed.String(),
		lg.TxHash,
		lg.BlockNumber,
	)
}
