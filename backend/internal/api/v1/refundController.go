package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"web3crowdfunding/internal/ethereum"
)

func StartRefundController(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/campaigns/refund/{campaignId}", refund)
}

func refund(w http.ResponseWriter, r *http.Request) {
	campaignId := r.PathValue("campaignId")

	campaignIdConverted, err := strconv.Atoi(campaignId)
	if err != nil {
		log.Println("Bad request: could not convert campaignId", campaignId, "error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transaction, err := ethereum.BuildRefundTransaction(campaignIdConverted)
	if err != nil {
		log.Println("Error when building transaction:", err)
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
		log.Println("Error writing response:", err)
		return
	}
}
