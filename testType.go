package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var typeTestKeyboard = map[string]tgbotapi.ReplyKeyboardMarkup{
	"anxiety":            anxietyKeyboard,
	"depression":         depressionKeyboard,
	"motivation":         motivationKeyboard,
	"communication":      communicationKeyboard,
	"affective":          affectiveKeyboard,
	"hamilton":           hamiltonSymptomsKeyboard,
	"beckanx":            beckAnxSymptomsKeyboard,
	"shihan":             shihanSymptomsKeyboard,
	"socialanx":          socialanxSymptomsKeyboard,
	"zung":               zungSymptomsKeyboard,
	"teylor":             yesNoSymptomsKeyboard,
	"geriatric":          yesNoSymptomsKeyboard,
	"motivationgoal":     yesNoSymptomsKeyboard,
	"motivationapproval": yesNoSymptomsKeyboard,
	"communicatelevel":   yesNoSymptomsKeyboard,
	"selfcontrol":        selfcontrolSymptomsKeyboard,
	"alexitimia":         alexitimiaSymptomsKeyboard,
	"autism":             autismSymptomsKeyboard,
}
