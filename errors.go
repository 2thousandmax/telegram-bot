package main

import (
	"errors"
)

var (
	invalidGroupError = errors.New("Неизвестная группа")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case invalidGroupError:
		messageText = b.messages.Errors.InvalidGroup
	default:
		messageText = b.messages.Errors.Default
	}

	b.SendTextMessage(chatID, messageText)
}
