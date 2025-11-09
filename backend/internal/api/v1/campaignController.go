package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/ethereum"
)

func StartController() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/campaign/create", create)
	http.HandleFunc("/campaign/create-unsigned", createUnsigned)
	http.HandleFunc("/campaign/donate", donate)
	http.HandleFunc("/campaign/get", get)
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

	unsignedCampaign, err := ethereum.CreateUnsignedCampaign(campaign)
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

	transaction, err := ethereum.CreateCampaign(campaign.Owner, campaign.Title, campaign.Description, &campaign.Target, &campaign.Deadline, campaign.Image)

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

func get(w http.ResponseWriter, r *http.Request) {
	campaigns, err := ethereum.GetCampaigns()

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

func donate(w http.ResponseWriter, r *http.Request) {
	var donation dtos.DonationDTO
	json.NewDecoder(r.Body).Decode(&donation)

	transaction, err := ethereum.DonateToCampaign(donation.CampaignId, donation.Value)

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
