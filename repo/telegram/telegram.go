package telegram

import (
	"bugsnag-notification/entity"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramRepo struct {
}

func (t TelegramRepo) PushToChannel(message entity.Payload) error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN_BOT_TELEGRAM"))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	text := "Event in " + message.Error.ReleaseStage + " from " + message.Account.Name + " in " + message.Error.StackTrace[0].File + " (details)\n" +
		"Unhandled error\n" +
		message.Error.ExceptionClass + " : " + message.Error.Message + "\n" +
		"Location\n" +
		message.Error.StackTrace[0].File + ":" + message.Error.StackTrace[0].LineNumber + "\n" +
		"URL : " + message.Error.URL
	msg := tgbotapi.NewMessageToChannel(os.Getenv("USERNAME_CHANNEL"), text)
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("error send message", err)
		return err
	}
	return nil
}
