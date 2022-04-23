package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Шихана"),
		tgbotapi.NewKeyboardButton("Социофобии"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Тейлора"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var depressionKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Занга"),
		tgbotapi.NewKeyboardButton("Гериатрическая"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

/*------------------------------------------------------------------*/

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
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Слегка. Не слишком меня беспокоил"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Умеренно. Это было неприятно, но я мог это переносить"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Умеренно. Это было неприятно, но я мог это переносить"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var shihanSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Нет"),
		tgbotapi.NewKeyboardButton("Немного"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Умеренно"),
		tgbotapi.NewKeyboardButton("Сильно"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Крайне сильно"),
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var socialanxSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Нет"),
		tgbotapi.NewKeyboardButton("Скорее нет"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Скорее да"),
		tgbotapi.NewKeyboardButton("Да"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var teylorSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Нет"),
		tgbotapi.NewKeyboardButton("Да"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var zungSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Никогда или изредка"),
		tgbotapi.NewKeyboardButton("Иногда"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Часто"),
		tgbotapi.NewKeyboardButton("Почти всегда или постоянно"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)

var geriatricSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Нет"),
		tgbotapi.NewKeyboardButton("Да"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться к началу"),
	),
)
