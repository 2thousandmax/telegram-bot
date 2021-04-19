package main

import (
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

// IsEvenWeek check is current week even
func IsEvenWeek(now time.Time) (int, string) {
	if _, thisWeek := now.ISOWeek(); thisWeek%10 != 0 {

		return 1, "Знаменатель"
	}
	return 0, "Числитель"
}

// ComposeMessage prepare text message
func ComposeMessage(group string, date time.Time, data Data) string {
	weekDay := strings.ToLower(date.Weekday().String())

	lesonsTable := data.Timetable[group][weekDay]
	weekTypeInt, weekTypeStr := IsEvenWeek(time.Now())

	msgText := fmt.Sprintf("Расписание для группы *%s*\n", group)
	msgText += fmt.Sprintf("Неделя: *%v*\n", weekTypeStr)
	msgText += fmt.Sprintf("День недели: *%v*\n\n", weekDayRu(weekDay))

	if weekDay == "sunday" {
		msgText += "Chill out! Пар нет"
		return msgText
	}

	for i := range lesonsTable {
		classes := lesonsTable[i][0]
		if len(lesonsTable[i]) > 1 {
			classes = lesonsTable[i][weekTypeInt]
		}

		msgText += fmt.Sprintf("*%v. %s*\n", i+1, classes[0])
		msgText += fmt.Sprintf("├ Время: *%v*\n", data.Rings[i])
		msgText += fmt.Sprintf("├ %s: *%v*\n", classes[1], classes[2])
		msgText += fmt.Sprintf("└ %s\n\n", classes[3])
	}

	return msgText
}

// weekDayRu translate weekday to russian equivalent
func weekDayRu(weekDay string) string {
	var text string

	switch weekDay {
	case "monday":
		text = "Понедельник"
	case "tuesday":
		text = "Вторник"
	case "wednesday":
		text = "Среда"
	case "thursday":
		text = "Четверг"
	case "friday":
		text = "Пятница"
	case "saturday":
		text = "Суббота"
	case "sunday":
		text = "Воскресенье"
	}

	return text
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
