package main

import (
	"errors"
)

var (
	ErrInvalidGroup    = errors.New("invalid group")
	ErrGroupNotFound   = errors.New("group not found")
	ErrInternalError   = errors.New("internal error")
	ErrMessageOutdated = errors.New("message is out of date")
	ErrUnknownCommand  = errors.New("unknown command")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	// switch err {
	// case ErrInvalidGroup:
	// 	messageText = b.messages.Errors.InvalidGroup
	// case ErrGroupNotFound:
	// 	messageText = b.messages.Errors.GroupNotFound
	// case ErrMessageOutdated:
	// 	messageText = b.messages.Errors.MessageOutdated
	// default:
	// 	messageText = b.messages.Errors.Default
	// }
	messageText = b.Config.Errors["default"]

	b.SendTextMessage(chatID, messageText, replyKeyboard)
}
