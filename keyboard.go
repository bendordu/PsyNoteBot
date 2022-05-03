package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var toStart = tgbotapi.NewKeyboardButton("Выход 🚪")

var startKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Пройти тест"),
		toStart,
	),
)

var backKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		toStart,
	),
)

var testKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Тревога"),
		tgbotapi.NewKeyboardButton("Депрессия"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Зависимость"),
		tgbotapi.NewKeyboardButton("Мотивация"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Коммуникация"),
		tgbotapi.NewKeyboardButton("Эмоции, настроение"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Прошлые результаты 🕔"),
	),
)

/*-----------------------------------------------------*/

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
		toStart,
	),
)

var depressionKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Занга"),
		tgbotapi.NewKeyboardButton("Гериатрическая"),
	),
	tgbotapi.NewKeyboardButtonRow(
		toStart,
	),
)

var motivationKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Потребность в достижении"),
		tgbotapi.NewKeyboardButton("Мотивации одобрения"),
	),
	tgbotapi.NewKeyboardButtonRow(
		toStart,
	),
)

var communicationKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Уровень общительности"),
		tgbotapi.NewKeyboardButton("Самоконтроль в общении"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Аутизм"),
		toStart,
	),
)

var affectiveKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Алекситимии"),
	),
	tgbotapi.NewKeyboardButtonRow(
		toStart,
	),
)

var addictionKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Кимберли-Янг"),
	),
	tgbotapi.NewKeyboardButtonRow(
		toStart,
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
		toStart,
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
		toStart,
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
		toStart,
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
		toStart,
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
		toStart,
	),
)

var selfcontrolSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Верно"),
		tgbotapi.NewKeyboardButton("Неверно"),
	),
	tgbotapi.NewKeyboardButtonRow(
		toStart,
	),
)

var alexitimiaSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Совершенно не согласен"),
		tgbotapi.NewKeyboardButton("Скорее не согласен"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Ни то, ни другое"),
		tgbotapi.NewKeyboardButton("Скорее согласен"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Совершенно согласен"),
		toStart,
	),
)

var autismSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Полностью согласен"),
		tgbotapi.NewKeyboardButton("Скорее согласен"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Скорее не согласен"),
		tgbotapi.NewKeyboardButton("Полностью не согласен"),
	),
	tgbotapi.NewKeyboardButtonRow(
		toStart,
	),
)

var yesNoSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Нет"),
		tgbotapi.NewKeyboardButton("Да"),
	),
	tgbotapi.NewKeyboardButtonRow(
		toStart,
	),
)

var kimberlyYangSymptomsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Никогда"),
		tgbotapi.NewKeyboardButton("Редко"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Регулярно"),
		tgbotapi.NewKeyboardButton("Часто"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Постоянно"),
		toStart,
	),
)
