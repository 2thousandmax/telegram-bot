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
	callbackYsterday = "Вчера"
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
		return b.handleRaspMessage(message)
	case messageSubjects:
		return b.handleSubjectsMessage(message)
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
		// return b.handleCallbackUnknown()
		return ErrUnknownCommand
	}
}


