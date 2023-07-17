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

type PasswordsOptions struct {
	CountChar      int
	CountPasswords int
	WithSymbols    bool
}

func basePasswordOptions() PasswordsOptions {
	return PasswordsOptions{
		CountChar:      8,
		CountPasswords: 5,
		WithSymbols:    false,
	}
}

func genPasswords(options PasswordsOptions) []string {
	var pswdArr []string

	for i := 1; i <= options.CountPasswords; i++ {
		password := make([]rune, options.CountChar)

		for c := range password {
			if options.WithSymbols {
				password[c] = lettersExtra[rand.Intn(len(letters))]
			} else {
				password[c] = letters[rand.Intn(len(letters))]
			}
		}

		pswdArr = append(pswdArr, string(password))
	}

	return pswdArr
}

func GetPswd() string {
	pswdArr := genPasswords(basePasswordOptions())

	var pswdHtml string = "Список сгенерированных паролей:\n"

	for i, password := range pswdArr {
		pswdHtml += strconv.Itoa(i+1) + ". <code>" + string(password) + "</code>\n"
	}

	return pswdHtml
}

func SetPswdSettings(update tgbotapi.Update) tgbotapi.MessageConfig {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите тип генерации:")
	msg.ReplyMarkup = numericKeyboard

	return msg
}
