package menu

import (
	"Qtools/app/controllers"
	"Qtools/app/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func handlePswdMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {
	switch text {
	case "pswd1":
		msg := tgbotapi.NewMessage(chatID, controllers.GetPswd())
		msg.ParseMode = "HTML"
		state[chatID] = StateMainMenu
		bot.Send(msg)
	case "pswd2":
		showPswdMenuSettings(bot, chatID)
	case "pswd3":
		showPswdMenuSettings1(bot, chatID)
	case "< back":
		state[chatID] = StateMainMenu
		showMainMenu(bot, chatID)
	}
}

func showPswdMenu(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "#️⃣ "+helpers.Bold("pswd")+" - a tool for generating passwords quickly")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Simple", "pswd1"),
			//tgbotapi.NewInlineKeyboardButtonData("Настраиваемый", "pswd2"),
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
