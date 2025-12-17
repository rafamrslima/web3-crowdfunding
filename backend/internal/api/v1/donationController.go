package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	dtos "web3crowdfunding/internal/DTOs"
	"web3crowdfunding/internal/db"
	"web3crowdfunding/internal/ethereum"
	"web3crowdfunding/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func StartDonationController(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/donations/adm/create", donate)
	mux.HandleFunc("/api/v1/donations/create", donateUnsigned)
	mux.HandleFunc("/api/v1/donations/{donor}", getDonationsByDonor)
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

		http.Error(w, err.Error(), http.StatusBadRequest)
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
	if err := json.NewDecoder(r.Body).Decode(&donation); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

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

func getDonationsByDonor(w http.ResponseWriter, r *http.Request) {
	donorId := r.PathValue("donor")
	if !common.IsHexAddress(donorId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid donor address"))
		return
	}
	addr := common.HexToAddress(donorId)
	donations, err := db.GetDonationsByDonor(addr.Bytes())

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(donations)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println("Error writing response:", err)
		return
	}
}
