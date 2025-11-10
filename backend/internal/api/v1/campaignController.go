package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/ethereum"
	"web3crowdfunding/internal/utils/validation"
)

func StartController() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/campaign/create", create)
	mux.HandleFunc("/campaign/create-unsigned", createUnsigned)
	mux.HandleFunc("/campaign/donate", donate)
	mux.HandleFunc("/campaign/get", get)

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

	err := validation.ValidateCampaign(campaign)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Bad Request: %v", err)
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

	err := validation.ValidateCampaign(campaign)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Bad Request: %v", err)
		return
	}

	transaction, err := ethereum.CreateCampaign(campaign)

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
