package clients

import (
	"go-telegram/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.Config("APP_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false
	return bot
}
