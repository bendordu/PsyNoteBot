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
		typesTest TypesTest
		psyParams = PsyParams{bot: bot}
	)
	testD := make(map[int64]TestData)

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

		InsertUser(userbot, db) ///////если нет - не вызывать

		if text == "/start" || text == "Вернуться к началу" { //Нулевой уровень 0 - старт

			userbot.Score, userbot.Number, userbot.Level, userbot.TestID = 0, 0, 0, 0
			UpdateData(userbot, db)

			psyParams.keyboard = testKeyboard
			psyParams.text = "Выберите тип"
			change(psyParams)

			setLevel(1, userbot, db)

		} else if text == "Прошлые результаты" {

			psyParams.text = SelectOldResult(userbot, typesTest, db)
			change(psyParams)

		} else if Select("level", userbot, db) == 1 { //Переходим на следующий уровень 1 - выбор типа тестов

			psyParams.text = text
			typeTest(psyParams, typesTest)

			setLevel(2, userbot, db)

		} else if Select("level", userbot, db) == 2 { //Выбор шкалы - 2 уровень

			psyParams.text = text
			testData := testDetails(psyParams, typesTest)
			testD[update.Message.Chat.ID] = testData

			InsertTestID(userbot, testData.NameEng, db)

			setLevel(3, userbot, db)

		} else if Select("level", userbot, db) == 3 { //Подсчет баллов при каждом новом выборе

			if Select("number", userbot, db) != 0 {

				userbot.Score = countScore(testD, update.Message.Chat.ID, text, Select("number", userbot, db)) + Select("score", userbot, db)
				UpdateScore(userbot, db)
			}

			if Select("number", userbot, db) < len(testD[update.Message.Chat.ID].Questions) {

				numberQuestionTest(psyParams, testD, update.Message.Chat.ID, Select("number", userbot, db))

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

				psyParams.text = result(Select("score", userbot, db), testD, update.Message.Chat.ID)
				change(psyParams)

				userbot.Score, userbot.Number, userbot.Level, userbot.TestID = 0, 0, 0, 0
				UpdateData(userbot, db)
				delete(testD, update.Message.Chat.ID)
			}
		}
	}
}
