package menu

import (
	"Qtools/app/controllers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func showPswdMenu(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Pswd menu")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Быстрый", "pswd1"),
			tgbotapi.NewInlineKeyboardButtonData("Настраиваемый", "pswd2"),
		),
	)
	bot.Send(msg)
}

func showPswdMenuSettings(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Password generator settings")

	options := controllers.BasePasswordOptions()
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Number of characters - "+strconv.Itoa(options.CountChar), "pswd3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Number of passwords - "+strconv.Itoa(options.CountPasswords), "pswd4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Use special characters - "+strconv.FormatBool(options.WithSymbols), "pswd5"),
		),
	)
	bot.Send(msg)
}

func showPswdMenuSettings1(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Enter new number of characters: ")
	bot.Send(msg)
}
