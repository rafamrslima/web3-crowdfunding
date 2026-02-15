package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(addr string) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)

	StartCampaignController(mux)
	StartDonationController(mux)
	StartWithdrawController(mux)
	StartRefundController(mux)
	StartCategoriesController(mux)

	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      WithCORS(mux),
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Server listening on %s", s.httpServer.Addr)

	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down HTTP server...")
	return s.httpServer.Shutdown(ctx)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
		if allowedOrigin == "" {
			allowedOrigin = "http://localhost:3000"
		}
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
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
