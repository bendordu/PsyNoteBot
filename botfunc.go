package main

import (
	"strconv"

	"log"

	"fmt"

	"io/ioutil"

	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func numberQuestionTest(p PsyParams, testData TestData, i int) {
	p.text = fmt.Sprintf("%s. %s", strconv.Itoa(i+1), testData.Questions[i])
	p.keyboard = typeTestKeyboard[testData.NameEng]
	change(p)
}

func countScore(testData TestData, text string, number int) (score int) {

	inv := false
	if len(testData.Inverse) != 0 {
		for _, n := range testData.Inverse {
			if number == n {
				inv = true
			}
		}
	}

	for ind, value := range testData.PointText {
		if text == value {
			if inv == false {
				score = testData.PointInt[ind]
			} else {
				score = testData.PointInt[len(testData.PointInt)-1-ind]
			}
		}
	}
	return score
}

func result(score int, testData TestData) (resultText string) {

	var resT string

	for ind, value := range testData.ResultSum {
		if score <= value && ind == 0 {
			resT = testData.ResultText[ind]

		} else if score <= testData.ResultSum[ind] && score > testData.ResultSum[ind-1] {
			resT = testData.ResultText[ind]
		}
	}
	resultText = fmt.Sprintf("Суммарное количество баллов: %s. %s", strconv.Itoa(score), resT)
	return resultText
}
