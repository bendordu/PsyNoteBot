package main

import (
	"strconv"

	"log"

	"io/ioutil"

	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func change(p PsyParams) {
	msg := tgbotapi.NewMessage(p.chatid, p.text)
	msg.ReplyMarkup = p.keyboard
	p.bot.Send(msg)
}

func typeTest(p PsyParams, typesTest TypesTest) {

	for _, typeTest := range typesTest.TestTypes {
		if typeTest.NameRus == p.text {
			p.text = "Выберите шкалу " + typeTest.Text
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

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &testData); err != nil {
		log.Fatal(err)
	}

	p.text = testData.TestAbout
	p.keyboard = startKeyboard
	change(p)
	return testData
}

func numberQuestionTest(p PsyParams, testData TestData, i int) {
	p.text = testData.TestQuestions[i]
	p.keyboard = typeTestKeyboard[testData.TestNameEng]
	change(p)
}

func countScore(testData TestData, text string) (score int) {
	for key, value := range testData.TestPoint {
		if text == key {
			score = value
		}
	}
	return score
}

func result(score int, testData TestData) (resultText string) {

	var resT string

	var keyArr []int
	for key := range testData.TestResult {
		keyArr = append(keyArr, key)
	}

	for ind, value := range keyArr {
		if (score <= value && ind == 0) || (score >= value && ind == len(keyArr)-1) {
			resT = testData.TestResult[value]

		} else if score <= keyArr[ind] && score > keyArr[ind-1] {
			resT = testData.TestResult[keyArr[ind]]
		}
	}
	resultText = "Суммарное количество баллов: " + strconv.Itoa(score) + ". " + resT
	return resultText
}
