package main

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	commandSetup = "setup"
	messageRasp = "Расписание"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandSetup:
		return b.handleSetupCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	// id := message.From.ID
	userName := message.From.UserName
	text := fmt.Sprintf(b.messages.Responses.Start, userName)

	return b.SendTextMessage(message.Chat.ID, text)
}

func (b *Bot) handleSetupCommand(message *tgbotapi.Message) error {
	return b.SendTextMessage(message.Chat.ID, b.messages.Responses.Setup)
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	return b.SendTextMessage(message.Chat.ID, b.messages.Responses.UnknownCommand)
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch message.Text {
	case messageRasp:
		return b.handleRaspMessage(message, b.data)
	default:
		// TODO separate function
		return b.SendTextMessage(message.Chat.ID, b.messages.Responses.UnknownCommand)
	}
}

func (b *Bot) handleRaspMessage(message *tgbotapi.Message, data Data) error {
	// TODO: get group from storage
	group := "IZ-21-1"
	weekDay := strings.ToLower(time.Now().Weekday().String())
	lesonsTable := data.Timetable[group][weekDay]
	weekType := IsEvenWeek(time.Now())

	mainTemplate := fmt.Sprintf("*Расписание для группы %s*\n", group)
	mainTemplate += fmt.Sprintf("Неделя: %v\n", weekType)
	mainTemplate += fmt.Sprintf("День недели: %v\n\n", weekDay)

	for i := range lesonsTable {
		var classes [4]interface{}

		if len(lesonsTable[i]) > 1 {
			classes = lesonsTable[i][weekType]
		} else {
			classes = lesonsTable[i][0]
		}

		mainTemplate += fmt.Sprintf("*%v. %s\n*", i+1, classes[0])
		mainTemplate += fmt.Sprintf("│ %s: %v\n", classes[1], classes[2])
		mainTemplate += fmt.Sprintf("└ %s\n\n", classes[3])
	}

	return b.SendTextMessage(message.Chat.ID, mainTemplate)
}

// IsEvenWeek check is current week even
func IsEvenWeek(now time.Time) int {
	if _, thisWeek := now.ISOWeek(); thisWeek%10 != 0 {
		return 0
	}

	return 1
}
