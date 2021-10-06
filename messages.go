package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) handleRaspMessage(message *tgbotapi.Message) error {
	return b.SendTextMessage(message.Chat.ID, "Расписание", inlineKeyboardGroup)
}

func (b *Bot) handleSubjectsMessage(message *tgbotapi.Message) error {
	// var msgText string
	// for i, v := range data.Classes {
	// 	msgText += fmt.Sprintf("*%v*. %s*(%s)*\n", i+1, v, data.Controls[i])
	// }
	return b.SendTextMessage(message.Chat.ID, "Предметы", replyKeyboard)
}
