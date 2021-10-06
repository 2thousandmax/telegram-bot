package main

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

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

	group, err := b.Storage.GetGroup(id)
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

func (b *Bot) handleCallbackUnknown(callback *tgbotapi.CallbackQuery) error {
	return nil
}
