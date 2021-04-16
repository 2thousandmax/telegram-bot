package main

import (
	"fmt"
	"regexp"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart     = "start"
	messageRasp      = "Расписание"
	messageSubjects  = "Список предметов"
	callbackTomorrow = "Завтра"
	callbackBack     = "Назад"
)

// Handle new command
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return ErrUnknownCommand
	}
}

// Handle new message
func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch message.Text {
	case messageRasp:
		return b.handleRaspMessage(message, b.data)
	case messageSubjects:
		return b.handleSubjectsMessage(message, b.data)
	default:
		return ErrUnknownCommand
	}
}

// Handle new callback query
func (b *Bot) handleCallbackQuery(callback *tgbotapi.CallbackQuery) error {
	// timeout := time.Unix(int64(callback.Message.Date), 0).Add(20 * time.Minute)
	// now := time.Now()

	// if timeout.Unix() < now.Unix() {
	// 	return ErrMessageOutdated
	// }

	matched, err := regexp.MatchString("ИЗ-2[12]-[12]", callback.Data)
	if err != nil {
		return ErrInternalError
	}

	switch {
	case matched:
		return b.handleGroupCallback(callback)
	case (callback.Data == callbackTomorrow):
		return b.handleTomorrowCallback(callback)
	case (callback.Data == callbackBack):
		return b.handleBackCallback(callback)
	default:
		return ErrUnknownCommand
	}
}

// Handle group callback
func (b *Bot) handleGroupCallback(callback *tgbotapi.CallbackQuery) error {
	id := callback.From.ID
	group := callback.Data

	if err := b.storage.SaveGroup(id, group); err != nil {
		return ErrInternalError
	}

	date := time.Now()
	msgText := ComposeMessage(group, date, b.data)

	inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Завтра →", callbackTomorrow),
		),
	)

	return b.EditMessage(callback.Message.Chat.ID, callback.Message.MessageID, msgText, &inlineKeyboard)
}

// Handle `Завтра` callback
func (b *Bot) handleTomorrowCallback(callback *tgbotapi.CallbackQuery) error {
	// Check if message old
	group, err := b.storage.GetGroup(callback.From.ID)
	if err != nil {
		return ErrInternalError
	}

	if group == "" {
		return ErrMessageOutdated
	}

	date := time.Now().AddDate(0, 0, 1)
	msgText := ComposeMessage(group, date, b.data)

	inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("← Назад", callbackBack),
		),
	)

	return b.EditMessage(callback.Message.Chat.ID, callback.Message.MessageID, msgText, &inlineKeyboard) // TODO keyboard
}

// Handle `Назад` callback
func (b *Bot) handleBackCallback(callback *tgbotapi.CallbackQuery) error {
	id := callback.From.ID

	group, err := b.storage.GetGroup(id)
	if err != nil {
		return ErrInternalError
	}

	if group == "" {
		return ErrMessageOutdated
	}

	date := time.Now()
	msgText := ComposeMessage(group, date, b.data)

	inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Завтра →", callbackTomorrow),
		),
	)

	return b.EditMessage(callback.Message.Chat.ID, callback.Message.MessageID, msgText, &inlineKeyboard)
}

// Handle `/start` command
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	username := message.From.UserName

	text := fmt.Sprintf(b.messages.Responses.Start, username)

	return b.SendTextMessage(message.Chat.ID, text, replyKeyboard)
}

// Handle `Расписание` message
func (b *Bot) handleRaspMessage(message *tgbotapi.Message, data Data) error {
	msgText := "Расписание"

	return b.SendTextMessage(message.Chat.ID, msgText, inlineKeyboardGroup)
}

// Handle `Список предметов` message
func (b *Bot) handleSubjectsMessage(message *tgbotapi.Message, data Data) error {
	var msgText string
	for i, v := range data.Classes {
		msgText += fmt.Sprintf("*%v*. %s*(%s)*\n", i+1, v, data.Controls[i])
	}
	return b.SendTextMessage(message.Chat.ID, msgText, replyKeyboard)
}
