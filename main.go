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
	userD := make(map[int64]map[string]int)

	data := readFile("json/typeTest.json")
	if err := json.Unmarshal(data, &typesTest); err != nil {
		log.Fatal(err)
	}

	for update := range updates {

		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		psyParams.chatid = chatID
		text := update.Message.Text

		_, found := userD[chatID]
		if found == false {
			userData := make(map[string]int)
			userD[chatID] = userData
		}

		InsertUser(chatID, db)

		if text == "/start" || text == "Вернуться к началу" { //Нулевой уровень 0 - старт

			userD[chatID]["score"],
				userD[chatID]["number"],
				userD[chatID]["level"],
				userD[chatID]["testID"] = 0, 0, 0, 0

			psyParams.keyboard = testKeyboard
			psyParams.text = "Выберите тип"
			change(psyParams)

			userD[chatID]["level"] = 1

		} else if text == "Прошлые результаты" {

			psyParams.text = SelectOldResult(chatID, typesTest, db)
			change(psyParams)

		} else if userD[chatID]["level"] == 1 { //Переходим на следующий уровень 1 - выбор типа тестов

			psyParams.text = text
			typeTest(psyParams, typesTest)
			userD[chatID]["level"] = 2

		} else if userD[chatID]["level"] == 2 { //Выбор шкалы - 2 уровень

			psyParams.text = text
			testData := testDetails(psyParams, typesTest)
			testD[chatID] = testData

			userD[chatID]["testID"] = SelectTestID(testD[chatID].NameEng, db)
			userD[chatID]["level"] = 3

		} else if userD[chatID]["level"] == 3 { //Подсчет баллов при каждом новом выборе

			if userD[chatID]["number"] != 0 {
				userD[chatID]["score"] += countScore(testD, chatID, text, userD[chatID]["number"])
			}

			if userD[chatID]["number"] < len(testD[chatID].Questions) {
				numberQuestionTest(psyParams, testD, chatID, userD[chatID]["number"])
				userD[chatID]["number"] += 1

			} else {

				psyParams.keyboard = backKeyboard

				tresult := Tresult{
					Result: userD[chatID]["score"],
					Date:   time.Now(),
					ChatID: chatID,
					TestID: userD[chatID]["testID"]}
				InsertResult(tresult, db)

				psyParams.text = result(userD[chatID]["score"], testD, chatID)
				change(psyParams)

				delete(testD, chatID)
				delete(userD, chatID)
			}
		}
	}
}
