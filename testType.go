package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var typeTestKeyboard = map[string]tgbotapi.ReplyKeyboardMarkup{
	"anxiety":    anxietyKeyboard,
	"depression": depressionKeyboard,
	"hamilton":   hamiltonSymptomsKeyboard,
	"beckanx":    beckAnxSymptomsKeyboard,
	"shihan":     shihanSymptomsKeyboard,
	"socialanx":  socialanxSymptomsKeyboard,
	"zung":       zungSymptomsKeyboard,
	"teylor":     teylorSymptomsKeyboard,
	"geriatric":  geriatricSymptomsKeyboard,
}
