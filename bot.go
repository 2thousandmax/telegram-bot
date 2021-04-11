package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	data     Data
	messages Messages
	storage  *UserStorage
}

func NewTelegramBot(bot *tgbotapi.BotAPI, d Data, msg Messages, storage *UserStorage) *Bot {
	return &Bot{
		bot:      bot,
		data:     d,
		messages: msg,
		storage:  storage,
	}
}

func (b *Bot) Start() error {
	// LongPolling
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	// // WebHooks
	// url := os.Getenv("PUBLIC_URL")
	// port := os.Getenv("PORT")
	// log.Printf("PORT: %s\nURL: %s",port, url)

	// token := os.Getenv("TELEGRAM_BOT_TOKEN")

	// _, err := b.bot.SetWebhook(tgbotapi.NewWebhook(fmt.Sprintf("https://%s/%s", url, token)))
	// if err != nil {
	// 	log.Fatalf("Problem in setting Webhook %v", err)
	// }

	// updates := b.bot.ListenForWebhook("/" + token)

	// go http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	for update := range updates {
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

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Расписание"),
		tgbotapi.NewKeyboardButton("Список предметов"),
	),
)

func (b *Bot) SendTextMessage(id int64, text string) error {
	msg := tgbotapi.NewMessage(id, text)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "markdown"

	_, err := b.bot.Send(msg)
	return err
}
