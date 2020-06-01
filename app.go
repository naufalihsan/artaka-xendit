package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go/client"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRET_KEY")
	artakaClient := client.New(secretKey)

	banks, _ := artakaClient.VirtualAccount.GetAvailableBanks()
	fmt.Println(banks)

}
