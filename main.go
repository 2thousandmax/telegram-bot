package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	config, err := NewConfig("configs/config.yaml")
	if err != nil {
		log.Panic(err)
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	log.Printf("Token: %s", token)

	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Authentication error: %v", err)
	}

	botApi.Debug = true
	log.Printf("Authorized on account %s", botApi.Self.UserName)

	bot := NewTelegramBot(botApi, config.Data, config.Messages)

	if err := bot.Start(); err != nil {
		log.Fatalf("Bot error: %v", err)
	}
}
