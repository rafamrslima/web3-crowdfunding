package main

import (
	"fmt"
	"log"
	"net/http"
	v1 "web3crowdfunding/internal/api/v1"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello web3")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env.")
	}

	v1.StartController()
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
