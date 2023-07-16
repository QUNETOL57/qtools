package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"strconv"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var lettersExtra = append(letters, []rune("!@#$%^&*")...)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Быстрый", "pswd1"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Настраиваемый", "pswd2"),
	),
)

func GenPasswords(countChar int) string {

	var pswdHtml string = ""

	for i := 1; i <= 5; i++ {
		password := make([]rune, countChar)

		for c := range password {
			password[c] = letters[rand.Intn(len(letters))]
		}

		pswdHtml += strconv.Itoa(i) + ". <code>" + string(password) + "</code>\n"
	}

	return pswdHtml
}

func SetPswdSettings(update tgbotapi.Update) tgbotapi.MessageConfig {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите тип генерации:")
	msg.ReplyMarkup = numericKeyboard

	return msg
}
