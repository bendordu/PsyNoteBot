package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var typeTestKeyboard = map[string]tgbotapi.ReplyKeyboardMarkup{
	"anxiety":  anxietyKeyboard,
	"hamilton": hamiltonSymptomsKeyboard,
	"beckanx":  beckAnxSymptomsKeyboard,
}
