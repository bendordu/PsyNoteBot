package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var testKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Тревога"),
		tgbotapi.NewKeyboardButton("Депрессия"),
		tgbotapi.NewKeyboardButton("Интеллект"),
	),
)

var anxietyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Гамильтона"),
		tgbotapi.NewKeyboardButton("Бека"),
		tgbotapi.NewKeyboardButton("Спилбергера-Ханина"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var hamiltonSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Отсутствие"),
		tgbotapi.NewKeyboardButton("В слабой степени"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("В умеренной степени"),
		tgbotapi.NewKeyboardButton("В тяжёлой степени"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("В очень тяжёлой степени"),
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var beckAnxSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Совсем не беспокоил"),
		tgbotapi.NewKeyboardButton("Слегка. Не слишком меня беспокоил"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Умеренно. Это было неприятно, но я мог это переносить"),
		tgbotapi.NewKeyboardButton("Очень сильно. Я с трудом мог это выносить"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var startKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Пройти тест"),
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var backKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)
