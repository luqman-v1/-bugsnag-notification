package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramRepo struct {
}

func (t TelegramRepo) Update() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN_BOT_TELEGRAM"))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	msg := tgbotapi.NewMessageToChannel("@bugsnagNotification", "Asdasdhadjs")
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("error send message", err)
	}
}
