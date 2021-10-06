package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	BotApi  *tgbotapi.BotAPI
	Config  Config
	Storage *Storage
	Debug   bool
}

func NewTelegramBot(config Config, storage *Storage) *Bot {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	log.Println("Token:", token)

	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("new telgram api: %v", err)
	}

	botApi.Debug = true
	log.Println("Authorized on account:", botApi.Self.UserName)

	return &Bot{
		BotApi:  botApi,
		Storage: storage,
		Debug:   true,
	}
}

func(b *Bot) GetUpdates() (tgbotapi.UpdatesChannel, error) {

	switch b.Debug {
	case true:
		_, err := b.BotApi.RemoveWebhook()
		if err != nil{
			return nil, fmt.Errorf("BotApi.RemoveWebhook: %v", err)
		}

		// LongPolling
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := b.BotApi.GetUpdatesChan(u)
		if err != nil {
			return nil, fmt.Errorf("BotApi.GetUpdatesChan: %v", err)
		}

		return updates, nil

	default:
		// WebHooks
		url := os.Getenv("PUBLIC_URL")
		port := os.Getenv("PORT")
		log.Printf("PORT: %s\nURL: %s", port, url)

		token := os.Getenv("TELEGRAM_BOT_TOKEN")

		_, err := b.BotApi.SetWebhook(tgbotapi.NewWebhook(fmt.Sprintf("https://%s/%s", url, token)))
		if err != nil {
			return nil, fmt.Errorf("bot.SetWebhook: %v", err)
		}

		updates := b.BotApi.ListenForWebhook("/" + token)

		go http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

		return updates, nil
	}

}

func (b *Bot) Start() error {

	updates, err := b.GetUpdates()
	if err != nil{
		return fmt.Errorf("bot.GetUpdates: %v", err)
	}

	for update := range updates {
		// Handle callback queries
		if update.CallbackQuery != nil {
			if err := b.handleCallbackQuery(update.CallbackQuery); err != nil {
				b.handleError(update.CallbackQuery.Message.Chat.ID, err)
			}
		}

		// Ignore any non-Message Updates
		if update.Message == nil {
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}

			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

	return nil
}

func (b *Bot) SendTextMessage(id int64, text string, replyMarkup interface{}) error {
	msg := tgbotapi.NewMessage(id, text)
	msg.ParseMode = "markdown"
	msg.ReplyMarkup = replyMarkup

	_, err := b.BotApi.Send(msg)
	return err
}

func (b *Bot) EditMessage(chatID int64, msgID int, text string, inlineKeyboard *tgbotapi.InlineKeyboardMarkup) error {
	msg := tgbotapi.NewEditMessageText(chatID, msgID, text)
	msg.ParseMode = "markdown"
	msg.ReplyMarkup = inlineKeyboard

	_, err := b.BotApi.Send(msg)
	return err
}
