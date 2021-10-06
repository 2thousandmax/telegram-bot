package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	username := message.From.UserName

	text := fmt.Sprintf(b.Config.Replies["default"], username)

	return b.SendTextMessage(message.Chat.ID, text, replyKeyboard)
}
