package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PsyTest struct {
	typeT, nameT string
	numberQ      int //Всего воспросов в тесте
	questions    []string
	pointTest    map[string]int
	keyboard     tgbotapi.ReplyKeyboardMarkup
	resultTest   map[int]string
}

type PsyParams struct {
	chatid   int64
	text     string
	keyboard tgbotapi.ReplyKeyboardMarkup
	bot      *tgbotapi.BotAPI
}

/*type PsyParamsChange struct {
	text     string
	keyboard tgbotapi.ReplyKeyboardMarkup
}

var Level1Params = PsyParamsChange{"Выберите тип", testKeyboard}*/
