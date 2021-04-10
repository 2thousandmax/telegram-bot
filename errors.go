package main

import (
	"errors"
)

var (
	errorInvalidGroup = errors.New("groupError")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case errorInvalidGroup:
		messageText = b.messages.Errors.InvalidGroup
	default:
		messageText = b.messages.Errors.Default
	}

	b.SendTextMessage(chatID, messageText)
}
