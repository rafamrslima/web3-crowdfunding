package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	v1 "web3crowdfunding/internal/api/v1"
	"web3crowdfunding/internal/indexer"
	"web3crowdfunding/internal/repositories"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello web3")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env.")
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer stop()

	httpServer := v1.NewServer(":8081")
	_, err := repositories.GetDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	go func() {
		log.Println("Starting HTTP server on :8081...")
		if err := httpServer.Start(); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	var listenerDone sync.WaitGroup
	listenerDone.Add(1)
	go func() {
		defer listenerDone.Done()
		log.Println("Starting chain event listener...")
		indexer.StartEventListener(ctx)
		log.Println("Event listener stopped")
	}()

	<-ctx.Done()
	log.Println("\nShutdown signal received...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	listenerDone.Wait()
	repositories.CloseDB()

	log.Println("Application stopped gracefully")
}
