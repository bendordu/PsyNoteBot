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
	TtestNameRus,
	TestNameEng,
	TestAbout string
	TestQuestions     []string
	QuantityQuestions int
	TestPoint         map[string]int
	TestResult        map[int]string
}
