package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart    = "start"
	commandSetup    = "setgroup"
	messageRasp     = "Расписание"
	messageSubjects = "Список предметов"
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
	id := message.From.ID
	username := message.From.UserName

	if err := b.storage.Save(id, username); err != nil {
		return err
	}

	text := fmt.Sprintf(b.messages.Responses.Start, username)

	return b.SendTextMessage(message.Chat.ID, text)
}

// Handle `/setup` command
func (b *Bot) handleSetupCommand(message *tgbotapi.Message) error {
	id := message.From.ID
	group := strings.Split(message.Text, " ")
	if len(group) < 2 {
		return errorInvalidGroup
	}

	matched, _ := regexp.MatchString(`ИЗ-\d\d-\d`, group[1])
	if !matched {
		return errorInvalidGroup
	}

	if err := b.storage.SetGroup(id, group[1], UsersBucket); err != nil {
		return err
	}

	msgText := fmt.Sprintf(b.messages.Responses.Setup, group[1])

	return b.SendTextMessage(message.Chat.ID, msgText)
}

// Handle unknown command
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	return b.SendTextMessage(message.Chat.ID, b.messages.Responses.UnknownCommand)
}

// Handle new message
func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch message.Text {
	case messageRasp:
		return b.handleRaspMessage(message, b.data)
	case messageSubjects:
		return b.handleSubjectsMessage(message, b.data)
	default:
		// TODO separate function
		return b.SendTextMessage(message.Chat.ID, b.messages.Responses.UnknownCommand)
	}
}

// Handle `Расписание` message
func (b *Bot) handleRaspMessage(message *tgbotapi.Message, data Data) error {
	id := message.From.ID
	group, err := b.storage.GetGroup(id, UsersBucket)
	if err != nil {
		// log.Fatal(err)
		return err
	}

	weekDay := strings.ToLower(time.Now().Weekday().String())
	lesonsTable := data.Timetable[group][weekDay]
	weekTypeInt, weekTypeStr := IsEvenWeek(time.Now())

	mainTemplate := fmt.Sprintf("*Расписание для группы %s*\n", group)
	mainTemplate += fmt.Sprintf("Неделя: %v\n", weekTypeStr)
	mainTemplate += fmt.Sprintf("День недели: %v\n\n", weekDayRu(weekDay))

	for i := range lesonsTable {
		classes := lesonsTable[i][0]
		if len(lesonsTable[i]) > 1 {
			classes = lesonsTable[i][weekTypeInt]
		}

		if weekDay != "sunday" {
			mainTemplate += fmt.Sprintf("*%v. %s\n*", i+1, classes[0])
			mainTemplate += fmt.Sprintf("│ %s: %v\n", classes[1], classes[2])
			mainTemplate += fmt.Sprintf("└ %s\n\n", classes[3])
		} else {
			mainTemplate += "Chill out! Пар нет"
		}

	}

	return b.SendTextMessage(message.Chat.ID, mainTemplate)
}

// Handle `Список предметов` message
func (b *Bot) handleSubjectsMessage(message *tgbotapi.Message, data Data) error {
	var msgText string
	for i, v := range data.Classes {
		msgText += fmt.Sprintf("%v. %s\n", i+1, v)
		// msgText += fmt.Sprintf("%v. %s\n └ %s\n", i+1, v, data.Lecturers[i])
	}
	return b.SendTextMessage(message.Chat.ID, msgText)
}
