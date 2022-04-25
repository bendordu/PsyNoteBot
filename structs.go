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
}

type TestBot struct {
	ID   int
	Name string
}

type Tresult struct {
	ID     int
	TestID int
	Result int
	Date   time.Time
}

type UserBot struct {
	ID        int
	ChatID    int64
	Level     int
	Score     int
	Number    int
	TresultID int
}
