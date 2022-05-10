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
	timeAnsw := make(map[int64]map[string]time.Time)

	data := readFile("json/typeTest.json")
	if err := json.Unmarshal(data, &typesTest); err != nil {
		log.Fatal(err)
	}

	for update := range updates {

		var (
			text   string
			chatID int64
		)

		if update.Message != nil {
			text = update.Message.Text
			chatID = update.Message.Chat.ID

		} else if update.CallbackQuery != nil {
			text = update.CallbackQuery.Data
			chatID = update.CallbackQuery.Message.Chat.ID

		} else {
			continue
		}

		psyParams.chatid = chatID

		_, found := userD[chatID]
		if found == false {
			userData := make(map[string]int)
			userD[chatID] = userData
		}

		_, found = timeAnsw[chatID]
		if found == false {
			timeAnswers := make(map[string]time.Time)
			timeAnsw[chatID] = timeAnswers
		}

		InsertUser(chatID, db)

		if text == "/start" || text == "Выход 🚪" { //Нулевой уровень 0 - старт

			if text == "/start" {
				psyParams.text = "👋🏼 Здравствуйте!\nЗдесь собраны психологические тесты, которые помогут определить ваше психическое состояние."
				change(psyParams)
			}

			setZero(chatID, userD)

			psyParams.keyboard = testKeyboard
			psyParams.text = "Выберите тип теста."
			change(psyParams)

			userD[chatID]["level"] = 1

		} else if text == "Прошлые результаты 🕔" || text == "/results" {

			str := SelectOldResult(chatID, typesTest, db)
			psyParams.keyboard = backKeyboard

			if len(str) == 0 {
				psyParams.text = "🦥 Пока результатов нет.\nПройдите любой из тестов до конца, и тогда они отобразятся здесь."

			} else {
				psyParams.text = "📅<b>Ваши результаты:</b>\n\n" + SelectOldResult(chatID, typesTest, db)
			}
			change(psyParams)

		} else if text == "/tests" {

			setZero(chatID, userD)
			userD[chatID]["level"] = 2
			psyParams.keyboard = backKeyboard
			psyParams.text = "⬇️"
			change(psyParams)
			allTests(chatID, typesTest, db, bot)

		} else if text == "/error" {

			setZero(chatID, userD)
			psyParams.text = "⚠️Если вы обнаружили ошибку, опишите ее в сообщении @stayclosetonight. Укажите, на каком этапе работы бота она возникла."
			psyParams.keyboard = backKeyboard
			change(psyParams)

		} else if text == "/description" {

			setZero(chatID, userD)
			psyParams.text = "\n\n❕Помните, что онлайн-тесты не предназначены для самостоятельной постановки диагноза. В случае любых сомнений обратитесь к специалисту.\n\nДля навигации по боту используйте команды из меню.\n\nЕсли вам интересна тема психического здоровья, подписывайтесь на <a href=\"https://t.me/parrhesia_psy\">канал</a>"
			psyParams.keyboard = backKeyboard
			change(psyParams)

		} else if text == "/author" {

			setZero(chatID, userD)
			psyParams.text = "👩‍💻Меня зовут Дарья, я — создатель этого бота, психолог, психиатр. Вы всегда можете обратиться ко мне, чтобы получить консультацию специалиста. Пишите: @stayclosetonight\n\n❤️Чтобы поблагодарить меня, вы можете оставить отзыв на <a href=\"https://www.b17.ru/daryadudinaa/\">сайте специалистов</a> или отправить пожертвования на <a href=\"https://yoomoney.ru/to/4100117806595904\">развитие проекта</a>.\n\nЕсли вам интересна тема психического здоровья, подписывайтесь на <a href=\"https://t.me/parrhesia_psy\">канал</a>"
			psyParams.keyboard = backKeyboard
			change(psyParams)

		} else if userD[chatID]["level"] == 1 { //Переходим на следующий уровень 1 - выбор типа тестов

			psyParams.text = text
			err := typeTest(psyParams, typesTest)
			if err != nil {
				log.Println(err)
			}
			userD[chatID]["level"] = 2

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

		} else if userD[chatID]["level"] == 3 || text == "Я понял. Продолжить тест" { //Подсчет баллов при каждом новом выборе

			/*else {
				if len(testD[chatID].Scales) != 0 {
					ans := make(map[int]int)
					answers[chatID] = ans
				}
			}*/
			timeAnsw[chatID]["timeUserAnsw"] = update.Message.Time()
			if timeAnsw[chatID]["timeUserAnsw"].Before(timeAnsw[chatID]["timeBotAnsw"].Round(time.Second)) {

				log.Println("Введено несколько ответов подряд на один вопрос")
				psyParams.text = "❗Не торопитесь. Еще раз прочитайте вопрос и ответьте заново."
				psyParams.keyboard = errKeyboard
				change(psyParams)
				userD[chatID]["number"]--

			} else {

				var (
					score int
					err   error
				)

				if userD[chatID]["number"] != 0 {

					if text != "Я понял. Продолжить тест" {
						score, err = countScore(testD, chatID, text, userD[chatID]["number"])
						userD[chatID]["score"] += score

						if err != nil {
							log.Println(err)
							psyParams.text = "❗Выберите вариант ответа, нажав кнопку на клавиатуре."
							psyParams.keyboard = typeTestKeyboard[testD[chatID].NameEng]
							change(psyParams)
						}
					}

				} else if text != "Пройти тест" {
					psyParams.text = "❗Не вводите ничего с клавиатуры. Выбирайте из предложенных вариантов ответов. Тест начался."
					change(psyParams)
				}

				if userD[chatID]["number"] < len(testD[chatID].Questions) {

					if err == nil {
						timeAnsw[chatID]["timeBotAnsw"] = time.Now()
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

					clearAll(chatID, testD, userD)
				}
			}
		}
	}
}

/*start - Вернуться к началу
description - Что может бот
tests - Вывести список тестов
results - Показать мои результаты
error - Сообщить об ошибке
author - О создателе бота*/

/*📑 психологические тесты по категориям:
-тревога
-коммуникация
-зависимости
-эмоции
-депрессия
-мотивация

Если вам интересна тема психического здоровья, подписывайтесь на <a href=\"https://t.me/parrhesia_psy\">канал</a>*/
