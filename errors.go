package main

import (
	"errors"
)

var (
	errorInvalidGroup = errors.New("groupError")
	errorGroupNotFound =  errors.New("groupNotFound")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case errorInvalidGroup:
		messageText = b.messages.Errors.InvalidGroup
	case errorGroupNotFound:
		messageText = b.messages.Errors.GroupNotFound
	default:
		messageText = b.messages.Errors.Default
	}

	b.SendTextMessage(chatID, messageText)
}
