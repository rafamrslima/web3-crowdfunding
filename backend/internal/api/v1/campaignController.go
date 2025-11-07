package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/ethereum"
)

func StartController() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/campaign/create", create)
	http.HandleFunc("/campaign/donate", donate)
	http.HandleFunc("/campaign/get", get)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func create(w http.ResponseWriter, r *http.Request) {
	var campaign dtos.CampaignDto
	json.NewDecoder(r.Body).Decode(&campaign)

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
