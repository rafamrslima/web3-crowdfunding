package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/database"
	"web3crowdfunding/internal/ethereum"
	"web3crowdfunding/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func StartCampaignController(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/campaigns", getAll)
	mux.HandleFunc("GET /api/v1/campaigns/onchain", getAllOnChain)
	mux.HandleFunc("GET /api/v1/campaigns/onchain/{id}", getById)
	mux.HandleFunc("GET /api/v1/campaigns/owner/{owner}", getCampaignsByOwner)
	mux.HandleFunc("POST /api/v1/campaigns/adm/create", create)
	mux.HandleFunc("POST /api/v1/campaigns/create", createUnsigned)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	campaigns, err := database.FetchAllCampaigns()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(campaigns)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func getAllOnChain(w http.ResponseWriter, r *http.Request) {
	campaigns, err := ethereum.FetchAllCampaigns()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(campaigns)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func getById(w http.ResponseWriter, r *http.Request) {
	campaignId := r.PathValue("id")

	if campaignId == "" {
		campaignId = "0"
	}

	campaignIdConverted, err := strconv.Atoi(campaignId)
	if err != nil {
		log.Println("Bad request:", campaignId)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	campaigns, err := ethereum.FetchCampaignById(campaignIdConverted)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(campaigns)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func getCampaignsByOwner(w http.ResponseWriter, r *http.Request) {
	ownerAddress := r.PathValue("owner")

	if ownerAddress == "" {
		log.Println("Bad request: invalid campaign owner")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addr := common.HexToAddress(ownerAddress)
	addrBytes := addr.Bytes()
	campaigns, err := database.GetCampaignsByOwner(addrBytes)

	if err != nil {
		log.Println("Error:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(campaigns)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var campaign dtos.CampaignDto
	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	err := utils.ValidateCampaign(campaign)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Bad Request: %v", err)
		return
	}

	transaction, err := ethereum.ExecuteCampaignCreation(campaign)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(data); err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func createUnsigned(w http.ResponseWriter, r *http.Request) {
	var campaign dtos.CampaignDto
	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	err := utils.ValidateCampaign(campaign)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Bad Request: %v", err)
		return
	}

	unsignedCampaign, err := ethereum.BuildCampaignTransaction(campaign)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(unsignedCampaign); err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}
