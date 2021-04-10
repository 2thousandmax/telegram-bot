package main

import (
	"encoding/binary"
	"time"
)

// IsEvenWeek check is current week even
func IsEvenWeek(now time.Time) (int, string) {
	if _, thisWeek := now.ISOWeek(); thisWeek%10 != 0 {
		return 0, "Знаменатель"
	}

	return 1, "Числитель"
}

// weekDayRu translate weekday to russian language
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
	case "sunday":
		text = "Суббота"
	case "saturday":
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
