package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	replyKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Расписание"),
			tgbotapi.NewKeyboardButton("Список предметов"),
		),
	)

	inlineKeyboardGroup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ИЗ-21-1", "ИЗ-21-1"),
			tgbotapi.NewInlineKeyboardButtonData("ИЗ-21-2", "ИЗ-21-2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ИЗ-22-1", "ИЗ-22-1"),
			tgbotapi.NewInlineKeyboardButtonData("ИЗ-22-2", "ИЗ-22-2"),
		),
	)

	inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Завтра →", callbackTomorrow),
		),
	)
)
