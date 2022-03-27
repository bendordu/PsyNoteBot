package main

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func change(chatid int64, text string, keyboard tgbotapi.ReplyKeyboardMarkup, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(chatid, text)
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func typeTest(chatid int64, text string, bot *tgbotapi.BotAPI) (typeT string) {

	switch text {
	case "Тревога":
		typeT = "anxiety"
		change(chatid, "Выберите шкалу тревоги", anxietyKeyboard, bot)

	case "Депрессия":
		typeT = "depression"
		change(chatid, "Выберите шкалу депрессии", anxietyKeyboard, bot) //!!
	}

	return typeT
}

func anxiety(chatid int64, text string, bot *tgbotapi.BotAPI) (nameT string, numberQ int) {
	switch text {
	case "Гамильтона":
		nameT = "hamilton"
		numberQ = 13
		change(chatid, aboutHamilton, hamiltonSymptomsKeyboard, bot)
		hamilton(chatid, bot, 0)
	}
	return nameT, numberQ
}

func hamilton(chatid int64, bot *tgbotapi.BotAPI, i int) {
	change(chatid, questionsHamilton[i], hamiltonSymptomsKeyboard, bot)
}

func psyTest(chatid int64, bot *tgbotapi.BotAPI, nameT string, numberQ int) {
	switch nameT {
	case "hamilton":
		hamilton(chatid, bot, numberQ)
	}
}

func result(score int, nameT string) (resultText string) {

	switch nameT {

	case "hamilton":
		switch {
		case score <= 17:
			resultText = strconv.Itoa(score) + " баллов. " + resultHamilton[17]
		case score <= 24 && score > 17:
			resultText = strconv.Itoa(score) + " баллов. " + resultHamilton[24]
		case score >= 25:
			resultText = strconv.Itoa(score) + " баллов. " + resultHamilton[25]
		}

	}
	return resultText
}
