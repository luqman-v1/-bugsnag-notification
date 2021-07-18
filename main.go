package main

import (
	"bugsnag-notification/gate"
	"bugsnag-notification/repo/telegram"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	telegram.TelegramRepo{}.Update()
	gate.Run()
}
