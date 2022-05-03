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
	answers := make(map[int64]map[int]int)

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

		if text == "/start" || text == "Выход 🚪" { //Нулевой уровень 0 - старт

			userD[chatID]["score"],
				userD[chatID]["number"],
				userD[chatID]["level"],
				userD[chatID]["testID"] = 0, 0, 0, 0

			psyParams.keyboard = testKeyboard
			psyParams.text = "Выберите тип"
			change(psyParams)

			userD[chatID]["level"] = 1

		} else if text == "Прошлые результаты 🕔" || text == "/show_results" {

			psyParams.text = SelectOldResult(chatID, typesTest, db)
			if len(psyParams.text) == 0 {
				psyParams.text = "🦥Пока результатов нет.\nПройдите любой из тестов до конца, и тогда они отобразятся здесь."
			}
			change(psyParams)

		} else if userD[chatID]["level"] == 1 { //Переходим на следующий уровень 1 - выбор типа тестов

			psyParams.text = text
			err := typeTest(psyParams, typesTest)
			if err != nil {
				log.Println(err)
			} else {
				userD[chatID]["level"] = 2
			}

		} else if userD[chatID]["level"] == 2 { //Выбор шкалы - 2 уровень

			psyParams.text = text
			testData, err := testDetails(psyParams, typesTest)
			if err != nil {
				log.Println(err)
			} else {
				testD[chatID] = testData
				userD[chatID]["testID"] = SelectTestID(testD[chatID].NameEng, db)
				userD[chatID]["level"] = 3
			}

		} else if userD[chatID]["level"] == 3 { //Подсчет баллов при каждом новом выборе

			/*else {
				if len(testD[chatID].Scales) != 0 {
					ans := make(map[int]int)
					answers[chatID] = ans
				}
			}*/

			var (
				score int
				err   error
			)
			if userD[chatID]["number"] != 0 {
				score, err = countScore(testD, chatID, text, userD[chatID]["number"])
				userD[chatID]["score"] += score
				if err != nil {
					log.Println(err)
					psyParams.text = "❗Выберите вариант ответа, нажав кнопку на клавиатуре."
					psyParams.keyboard = typeTestKeyboard[testD[chatID].NameEng]
					change(psyParams)
				}
			} else if text != "Пройти тест" {
				psyParams.text = "❗Не вводите ничего с клавиатуры. Выбирайте из предложенных вариантов ответов. Тест начался."
				change(psyParams)
			}

			if userD[chatID]["number"] < len(testD[chatID].Questions) {
				if err == nil {
					numberQuestionTest(psyParams, testD, chatID, userD[chatID]["number"])
					userD[chatID]["number"] += 1
				}

			} else {

				psyParams.keyboard = backKeyboard

				tresult := Tresult{
					Result: userD[chatID]["score"],
					Date:   time.Now(),
					ChatID: chatID,
					TestID: userD[chatID]["testID"]}
				InsertResult(tresult, db)

				psyParams.text = result(answers, userD[chatID]["score"], testD, chatID)
				change(psyParams)

				delete(testD, chatID)
				delete(userD, chatID)
			}
		}
	}
}
