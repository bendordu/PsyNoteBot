package main

import (
	"log"

	"time"

	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot := tbot()
	db := InitDB("db.sqlite3")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	var (
		testData  TestData
		typesTest TypesTest
		psyParams = PsyParams{bot: bot}
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
		userbot := UserBot{ChatID: update.Message.Chat.ID}

		InsertUser(userbot, db)

		if text == "/start" || text == "Вернуться к началу" { //Нулевой уровень 0 - старт

			userbot.Score, userbot.Number, userbot.Level, userbot.TestID = 0, 0, 0, 0
			UpdateData(userbot, db)

			psyParams.keyboard = testKeyboard
			psyParams.text = "Выберите тип"
			change(psyParams)

			userbot.Level = 1
			UpdateLevel(userbot, db)

		} else if Select("level", userbot, db) == 1 { //Переходим на следующий уровень 1 - выбор типа тестов

			psyParams.text = text
			typeTest(psyParams, typesTest)

			userbot.Level = 2
			UpdateLevel(userbot, db)

		} else if Select("level", userbot, db) == 2 { //Выбор шкалы - 2 уровень

			psyParams.text = text
			testData = testDetails(psyParams, typesTest)

			InsertTestID(userbot, testData.NameEng, db)

			userbot.Level = 3
			UpdateLevel(userbot, db)

		} else if Select("level", userbot, db) == 3 { //Подсчет баллов при каждом новом выборе

			if Select("number", userbot, db) != 0 {

				userbot.Score = countScore(testData, text, Select("number", userbot, db)) + Select("score", userbot, db)
				UpdateScore(userbot, db)
			}

			if Select("number", userbot, db) < len(testData.Questions) {

				numberQuestionTest(psyParams, testData, Select("number", userbot, db))

				userbot.Number = Select("number", userbot, db) + 1
				UpdateNumber(userbot, db)

			} else {

				psyParams.keyboard = backKeyboard

				tresult := Tresult{
					Result: Select("score", userbot, db),
					Date:   time.Now(),
					UserID: Select("user_id", userbot, db),
					TestID: Select("test_id", userbot, db)}
				InsertResult(tresult, db)

				psyParams.text = result(Select("score", userbot, db), testData)
				change(psyParams)

				userbot.Score, userbot.Number, userbot.Level, userbot.TestID = 0, 0, 0, 0
				UpdateData(userbot, db)
			}
		}
	}
}
