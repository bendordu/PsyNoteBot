package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5239188791:AAFKHNX1HqkSNJbbqu9amon8_Fm-v4Aws_w")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	var chatId int64
	var level int = 0
	var typeT, nameT string
	score := 0
	number := 0
	numberQ := 0

	for update := range updates {

		if update.Message == nil {
			continue
		}

		chatId = update.Message.Chat.ID
		text := update.Message.Text

		//Нулевой уровень 0 - старт
		if text == "/start" {
			level = 0
			change(chatId, "Выберите тип", testKeyboard, bot)
			level++

		} else if text == "Вернуться к началу" {
			level = 0
			change(chatId, "Выберите тип", testKeyboard, bot)
			level++

		} else if level == 1 { //Переходим на следующий уровень 1 - выбор типа тестов

			level++
			typeT = typeTest(chatId, text, bot)

		} else if level == 2 { //Выбор шкалы - 2 уровень

			level++
			switch typeT {
			case "anxiety":
				nameT, numberQ = anxiety(chatId, text, bot)
			}

		} else if level == 3 {

			for key, value := range point {
				if text == key {
					score += value
				}
			}
			if number < numberQ {
				number++
				psyTest(chatId, bot, nameT, number)
			} else {
				textResult := result(score, nameT)
				change(chatId, textResult, backKeyboard, bot)
				score = 0
				number = 0
				level = 0
			}

		}
	}
}
