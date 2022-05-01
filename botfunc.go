package main

import (
	"strconv"

	"log"

	"fmt"

	"io/ioutil"

	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func tbot() (bot *tgbotapi.BotAPI) {

	bot, err := tgbotapi.NewBotAPI(BOTAPI)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	return bot
}

func readFile(path string) (data []byte) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func change(p PsyParams) {
	msg := tgbotapi.NewMessage(p.chatid, p.text)
	msg.ReplyMarkup = p.keyboard
	msg.ParseMode = "HTML"
	p.bot.Send(msg)
}

func typeTest(p PsyParams, typesTest TypesTest) {

	for _, typeTest := range typesTest.TestTypes {
		if typeTest.NameRus == p.text {
			p.text = fmt.Sprintf("Выберите шкалу %s", typeTest.Text)
			p.keyboard = typeTestKeyboard[typeTest.NameEng]
		}
	}
	change(p)
}

func testDetails(p PsyParams, typesTest TypesTest) (testData TestData) {

	var path string
	for _, typeTest := range typesTest.TestTypes {
		for _, test := range typeTest.Tests {
			if test.TestNameRus == p.text {
				path = test.Path
			}
		}
	}

	data := readFile(path)
	if err := json.Unmarshal(data, &testData); err != nil {
		log.Fatal(err)
	}

	p.text = testData.About
	p.keyboard = startKeyboard
	change(p)
	return testData
}

func numberQuestionTest(p PsyParams, testD map[int64]TestData, chatID int64, i int) {
	p.text = fmt.Sprintf("%s. %s", strconv.Itoa(i+1), testD[chatID].Questions[i])
	p.keyboard = typeTestKeyboard[testD[chatID].NameEng]
	change(p)
}

func countScore(testD map[int64]TestData, chatID int64, text string, number int) (score int) {

	inv := false
	if len(testD[chatID].Inverse) != 0 {
		for _, n := range testD[chatID].Inverse {
			if number == n {
				inv = true
			}
		}
	}

	for ind, value := range testD[chatID].PointText {
		if text == value {
			if inv == false {
				score = testD[chatID].PointInt[ind]
			} else {
				score = testD[chatID].PointInt[len(testD[chatID].PointInt)-1-ind]
			}
		}
	}
	return score
}

func result(answers map[int64]map[int]int, score int, testD map[int64]TestData, chatID int64) (resultText string) {

	var resT string

	for ind, value := range testD[chatID].ResultSum {
		if score <= value && ind == 0 {
			resT = testD[chatID].ResultText[ind]

		} else if score <= testD[chatID].ResultSum[ind] && score > testD[chatID].ResultSum[ind-1] {
			resT = testD[chatID].ResultText[ind]
		}
	}

	if len(answers) != 0 {
		var totalSc int
		maxSc := make(map[string]int)
		for key, value := range testD[chatID].Scales {
			for _, numb := range value {
				totalSc += answers[chatID][numb]
			}
			maxSc[key] = totalSc
		}
	}

	resultText = fmt.Sprintf("Суммарное количество баллов: %s. %s", strconv.Itoa(score), resT)
	return resultText
}
