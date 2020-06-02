package main

import (
	"github.com/joho/godotenv"
	"github.com/naufalihsan/artaka-xendit/qris"
	"github.com/skip2/go-qrcode"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRET_KEY")
	client := qris.New(secretKey)

	createQRISData := qris.CreateQRISParams{
		ExternalID:  "c4n5ky-STud1o",
		Type:        qris.QRISTypeDYNAMIC,
		CallbackURL: "https://naufalihsan.co.id/callback",
		Amount:      99000,
	}
	resp, _ := client.QRIS.CreateQRIS(&createQRISData)

	err = qrcode.WriteFile(resp.QRString, qrcode.Medium, 256, "qris.png")
	if err != nil {
		log.Fatal("Error generate file")
	}
}
