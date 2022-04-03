package main

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func change(p PsyParams) {
	msg := tgbotapi.NewMessage(p.chatid, p.text)
	msg.ReplyMarkup = p.keyboard
	p.bot.Send(msg)
}

func typeTest(p PsyParams) (typeT string) {

	switch p.text {
	case "Тревога":
		typeT = "anxiety"
		p.text = "Выберите шкалу тревоги"
		p.keyboard = anxietyKeyboard

	case "Депрессия":
		typeT = "depression"
		p.text = "Выберите шкалу депрессии"
		p.keyboard = anxietyKeyboard //!!

	}
	change(p)
	return typeT
}

func anxiety(p PsyParams) (test PsyTest) {

	switch p.text {
	case "Гамильтона":
		test = Hamilton
		p.text = aboutHamilton

	case "Бека":
		test = BeckAnx
		p.text = aboutBeckAnx

	}
	p.keyboard = startKeyboard
	change(p)
	return test
}

func numberQuestionTest(p PsyParams, pT PsyTest, i int) {
	p.text = pT.questions[i]
	p.keyboard = pT.keyboard
	change(p)
}

func countScore(pT PsyTest, text string) (score int) {
	for key, value := range pT.pointTest {
		if text == key {
			score = value
		}
	}
	return score
}

func result(score int, pT PsyTest) (resultText string) {

	var resT string

	var keyArr []int
	for key := range pT.resultTest {
		keyArr = append(keyArr, key)
	}

	for ind, value := range keyArr {
		if (score <= value && ind == 0) || (score >= value && ind == len(keyArr)-1) {
			resT = resultHamilton[value]

		} else if score <= keyArr[ind] && score > keyArr[ind-1] {
			resT = resultHamilton[keyArr[ind]]
		}
	}
	resultText = "Суммарное количество баллов: " + strconv.Itoa(score) + ". " + resT
	return resultText
}
