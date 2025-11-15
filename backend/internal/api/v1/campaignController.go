package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/ethereum"
	"web3crowdfunding/internal/utils"
)

func StartController() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)

	mux.HandleFunc("/api/v1/campaigns", getAll)
	mux.HandleFunc("/api/v1/campaigns/{id}", getById)
	mux.HandleFunc("/api/v1/campaigns/create", create)
	mux.HandleFunc("/api/v1/campaigns/unsigned", createUnsigned)

	mux.HandleFunc("/api/v1/donations", donate)
	mux.HandleFunc("/api/v1/donations/unsigned", donateUnsigned)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", WithCORS(mux)))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
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

func getAll(w http.ResponseWriter, r *http.Request) {
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
		log.Println("Bad request:", campaignIdConverted)
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

func donateUnsigned(w http.ResponseWriter, r *http.Request) {
	var donation dtos.DonationDTO
	if err := json.NewDecoder(r.Body).Decode(&donation); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	err := utils.ValidateDonation(donation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Bad Request: %v", err)
		return
	}

	transaction, err := ethereum.BuildDonationTransaction(donation.CampaignId, donation.Value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(transaction)
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

func donate(w http.ResponseWriter, r *http.Request) {
	var donation dtos.DonationDTO
	json.NewDecoder(r.Body).Decode(&donation)

	transaction, err := ethereum.ExecuteDonationToCompaign(donation.CampaignId, donation.Value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(transaction)
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

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*") // or specific domain e.g. http://localhost:5173
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Respond to preflight (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
