package v1

import (
	"fmt"
	"log"
	"net/http"
)

func StartController() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)

	StartCampaignController(mux)
	StartDonationController(mux)
	StartWithdrawController(mux)
	StartRefundController(mux)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", WithCORS(mux)))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
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
