package main

import (
	"log"

	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(BOTAPI)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	level, score, number := 0, 0, 0 //Уровень выбора, Итоговый балл, Номер вопроса

	psyParams := PsyParams{bot: bot}
	var (
		testData  TestData
		typesTest TypesTest
	)

	data := readFile("json/typeTest.json")

	if err := json.Unmarshal(data, &typesTest); err != nil {
		log.Fatal(err)
	}

	for update := range updates {

		if update.Message == nil {
			continue
		}

		psyParams.chatid = update.Message.Chat.ID
		text := update.Message.Text

		//Нулевой уровень 0 - старт
		if text == "/start" || text == "Вернуться к началу" {
			score, number, level = 0, 0, 0
			psyParams.keyboard = testKeyboard
			psyParams.text = "Выберите тип"
			change(psyParams)
			level = 1

		} else if level == 1 { //Переходим на следующий уровень 1 - выбор типа тестов
			level = 2
			psyParams.text = text
			typeTest(psyParams, typesTest)

		} else if level == 2 { //Выбор шкалы - 2 уровень
			level = 3
			psyParams.text = text
			testData = testDetails(psyParams, typesTest)

		} else if level == 3 { //Подсчет баллов при каждом новом выборе

			if number != 0 {
				score += countScore(testData, text, number)
			}

			if number < len(testData.Questions) {
				numberQuestionTest(psyParams, testData, number)
				number++

			} else {
				psyParams.keyboard = backKeyboard
				psyParams.text = result(score, testData)
				change(psyParams)
				score, number, level = 0, 0, 0
			}

		}
	}
}
