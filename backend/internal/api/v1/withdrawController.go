package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"web3crowdfunding/internal/ethereum"
)

func StartWithdrawController(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/campaigns/withdraw/{id}", withdraw)
}

func withdraw(w http.ResponseWriter, r *http.Request) {
	campaignId := r.PathValue("id")

	campaignIdConverted, err := strconv.Atoi(campaignId)
	if err != nil {
		log.Println("Bad request: could not convert campaignId", campaignId, "error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transaction, err := ethereum.BuildWithdrawTransaction(campaignIdConverted)
	if err != nil {
		fmt.Println("Error when building transaction:", err)
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
