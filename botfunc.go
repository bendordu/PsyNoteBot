package main

import (
	"database/sql"
	"errors"
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

func clearAll(chatID int64, testD map[int64]TestData, userD map[int64]map[string]int) {
	delete(testD, chatID)
	delete(userD, chatID)
}

func setZero(chatID int64, userD map[int64]map[string]int) {
	userD[chatID]["score"],
		userD[chatID]["number"],
		userD[chatID]["level"],
		userD[chatID]["testID"] = 0, 0, 0, 0
}

func allTests(chatID int64, typesTest TypesTest, db *sql.DB, bot *tgbotapi.BotAPI) {
	testsList := SelectTests(typesTest, db)
	var testsNameRus [][]tgbotapi.InlineKeyboardButton

	msg := tgbotapi.NewMessage(chatID, "Список всех тестов")

	for _, test := range testsList {
		testsNameRus = append(testsNameRus, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(test, test)))
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(testsNameRus...)
	bot.Send(msg)
}

func typeTest(p PsyParams, typesTest TypesTest) (err error) {

	text := p.text
	for _, typeTest := range typesTest.TestTypes {
		if typeTest.NameRus == text {

			p.text = fmt.Sprintf("Выберите шкалу %s", typeTest.Text)
			p.keyboard = typeTestKeyboard[typeTest.NameEng]
			err = nil
			break

		} else {
			p.text = "❗Выберите тип теста, нажав кнопку на клавиатуре."
			err = errors.New("Не выбран тип теста")
		}
	}
	change(p)
	return err
}

func testDetails(p PsyParams, typesTest TypesTest) (testData TestData, err error) {

	text := p.text
out:
	for _, typeTest := range typesTest.TestTypes {
		for _, test := range typeTest.Tests {
			if test.TestNameRus == text {

				path := test.Path
				data := readFile(path)

				if err := json.Unmarshal(data, &testData); err != nil {
					log.Fatal(err)
				}

				p.text = testData.About
				p.keyboard = startKeyboard
				err = nil
				break out

			} else {
				p.text = "❗Выберите шкалу, нажав кнопку на клавиатуре."
				err = errors.New("Не выбрана шкала")
			}
		}
	}
	change(p)
	return testData, err
}

func numberQuestionTest(p PsyParams, testD map[int64]TestData, chatID int64, i int) {

	p.text = fmt.Sprintf("%s. %s", strconv.Itoa(i+1), testD[chatID].Questions[i])
	p.keyboard = typeTestKeyboard[testD[chatID].NameEng]
	change(p)
}

func countScore(testD map[int64]TestData, chatID int64, text string, number int) (score int, err error) {

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
			err = nil
			break

		} else {
			err = errors.New("Не выбран вариант ответа")
		}
	}
	return score, err
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

	resultText = fmt.Sprintf("Суммарное количество баллов: %s. %s\n\n❕Помните, что онлайн-тесты не предназначены для самостоятельной постановки диагноза. В случае любых сомнений обратитесь к специалисту. Например, вы можете написать создателю этого бота — психиатру @stayclosetonight", strconv.Itoa(score), resT)
	return resultText
}
