package main

import (
	"fmt"
	"log"
	v1 "web3crowdfunding/internal/api/v1"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello web3")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env.")
	}

	v1.StartController()
}
