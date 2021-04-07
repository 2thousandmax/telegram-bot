package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	data     Data
	messages Messages
}

func NewTelegramBot(bot *tgbotapi.BotAPI, d Data, msg Messages) *Bot {
	return &Bot{
		bot:      bot,
		data:     d,
		messages: msg,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

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
	),
)

func (b *Bot) SendTextMessage(id int64, text string) error {
	msg := tgbotapi.NewMessage(id, text)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "markdown"

	_, err := b.bot.Send(msg)
	return err
}
