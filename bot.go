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
	// b.bot.RemoveWebhook()

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
	// log.Printf("PORT: %s\nURL: %s", port, url)

	// token := os.Getenv("TELEGRAM_BOT_TOKEN")

	// _, err := b.bot.SetWebhook(tgbotapi.NewWebhook(fmt.Sprintf("https://%s/%s", url, token)))
	// if err != nil {
	// 	log.Fatalf("Problem in setting Webhook %v", err)
	// }

	// updates := b.bot.ListenForWebhook("/" + token)

	// go http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

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

var (
	replyKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Расписание"),
			tgbotapi.NewKeyboardButton("Список предметов"),
		),
	)

	inlineKeyboardGroup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ИЗ-21-1", "ИЗ-21-1"),
			tgbotapi.NewInlineKeyboardButtonData("ИЗ-21-2", "ИЗ-21-2"),
		),
	)

	inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Завтра →", callbackTomorrow),
		),
	)
)

func (b *Bot) SendTextMessage(id int64, text string, replyMarkup interface{}) error {
	msg := tgbotapi.NewMessage(id, text)
	msg.ParseMode = "markdown"
	msg.ReplyMarkup = replyMarkup

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) EditMessage(chatID int64, msgID int, text string, inlineKeyboard *tgbotapi.InlineKeyboardMarkup) error {
	msg := tgbotapi.NewEditMessageText(chatID, msgID, text)
	msg.ParseMode = "markdown"
	msg.ReplyMarkup = inlineKeyboard

	_, err := b.bot.Send(msg)
	return err
}
