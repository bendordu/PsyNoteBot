package main

import (
	"log"

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
	var typeT string                //Ключи, по которым осуществяется переход

	psyParams := PsyParams{bot: bot}
	var pTest PsyTest

	for update := range updates {

		if update.Message == nil {
			continue
		}

		psyParams.chatid = update.Message.Chat.ID
		text := update.Message.Text

		//Нулевой уровень 0 - старт
		if text == "/start" || text == "Вернуться к началу" {
			level = 0
			psyParams.keyboard = testKeyboard
			psyParams.text = "Выберите тип"
			change(psyParams)
			level++

		} else if level == 1 { //Переходим на следующий уровень 1 - выбор типа тестов
			level++
			psyParams.text = text
			typeT = typeTest(psyParams)

		} else if level == 2 { //Выбор шкалы - 2 уровень
			level++
			psyParams.text = text
			switch typeT {
			case "anxiety":
				pTest = anxiety(psyParams)
			}

		} else if level == 3 { //Подсчет баллов при каждом новом выборе

			score += countScore(pTest, text)

			if number < pTest.numberQ {
				numberQuestionTest(psyParams, pTest, number)
				number++

			} else {
				psyParams.keyboard = backKeyboard
				psyParams.text = result(score, pTest)
				change(psyParams)
				score, number, level = 0, 0, 0
			}

		}
	}
}
