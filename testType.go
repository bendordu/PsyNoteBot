package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var typeTestKeyboard = map[string]tgbotapi.ReplyKeyboardMarkup{
	"anxiety":            anxietyKeyboard,
	"depression":         depressionKeyboard,
	"motivation":         motivationKeyboard,
	"communication":      communicationKeyboard,
	"affective":          affectiveKeyboard,
	"addiction":          addictionKeyboard,
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
	"kimberlyYang":       kimberlyYangSymptomsKeyboard,
	"ptsr":               ptsrSymptomsKeyboard,
	"computeradd":        computeraddSymptomsKeyboard,
	"holland":            hollandSymptomsKeyboard,
	"eat26":              eat26SymptomsKeyboard,
	"dissociative":       dissociativeSymptomsKeyboard,
	"adhd":               adhdSymptomsKeyboard,
}
