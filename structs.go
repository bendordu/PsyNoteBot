package main

import (
	"time"

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
	Questions  []string
	PointText  []string
	PointInt   []int
	ResultText []string
	ResultSum  []int
	Inverse    []int
	Scales     map[string][]int
}

type Tresult struct {
	TestID int
	Result int
	Date   time.Time
	ChatID int64
}
