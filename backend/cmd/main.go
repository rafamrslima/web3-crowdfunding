package main

import (
	"fmt"
	"log"
	"sync"
	v1 "web3crowdfunding/internal/api/v1"
	"web3crowdfunding/internal/indexer"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello web3")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env.")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		v1.StartController()
	}()

	go func() {
		defer wg.Done()
		indexer.StartEventListener()
	}()

	wg.Wait()
}
