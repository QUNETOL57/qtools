package menu

import (
	"Qtools/app/controllers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func showMainMenu(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Main Menu")
	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/pswd"),
			tgbotapi.NewKeyboardButton("/tcase"),
			tgbotapi.NewKeyboardButton("/tarray"),
		),
	)
	msg.ReplyMarkup = replyKeyboard
	replyKeyboard.OneTimeKeyboard = true
	bot.Send(msg)
}

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

func showPswdMenuHard(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Pswd Menu Hard")

	options := controllers.BasePasswordOptions()
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Количество символов - "+strconv.Itoa(options.CountChar), "pswd3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Количество паролей - "+strconv.Itoa(options.CountPasswords), "pswd4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Использовать спец символы - "+strconv.FormatBool(options.WithSymbols), "pswd5"),
		),
	)
	bot.Send(msg)
}

func showTcaseMenu(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Tcase Введите текст: ")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("< back"),
		),
	)
	bot.Send(msg)
}

func showTarrayMenu(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Tarray")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("< back"),
		),
	)
	bot.Send(msg)
}
