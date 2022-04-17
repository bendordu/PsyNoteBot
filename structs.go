package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PsyParams struct {
	chatid   int64
	text     string
	keyboard tgbotapi.ReplyKeyboardMarkup
	bot      *tgbotapi.BotAPI
}

type Test struct {
	Path,
	TestName,
	TestNameRus string
}

type TypeTest struct {
	NameEng,
	NameRus,
	Text string
	Tests []Test
}

type TypesTest struct {
	TestTypes []TypeTest
}

type TestData struct {
	NameRus,
	NameEng,
	About string
	Questions []string
	Point     map[string]int
	Result    map[int]string
}
