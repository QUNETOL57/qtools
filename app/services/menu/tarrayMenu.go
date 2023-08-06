package menu

import (
	"Qtools/app/controllers"
	"Qtools/app/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var textToArray string
var quotesType string

func handleTarrayMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {
	switch text {
	case controllers.TArrayFristMenuStr:
		quotesType = controllers.TArrayWithoutQuotes
		showTarrayStyle(bot, chatID)
	case controllers.TArraySecondMenuStr:
		showTarraySecondMenu(bot, chatID)
	case controllers.TArraySingleQuotes, controllers.TArrayDoubleQuotes:
		quotesType = text
		showTarrayStyle(bot, chatID)
	case controllers.TArrayRound, controllers.TArraySquare, controllers.TArrayCurly, controllers.TArrayOld:
		sendArray(bot, chatID, text)
		state[chatID] = StateMainMenu
		showMainMenu(bot, chatID)
	case "< back":
		state[chatID] = StateMainMenu
		showMainMenu(bot, chatID)
	default:
		textToArray = text
		showTarrayFirstMenu(bot, chatID)
	}
}

func showTarrayText(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "#️⃣ "+helpers.Bold("tarray")+" - is a tool for quickly converting a series of values from a string to an array \n\n"+
		"Enter text separated by space or newline:")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("< back"),
		),
	)
	bot.Send(msg)
}

func showTarrayFirstMenu(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Choice generation type:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Simple", controllers.TArrayFristMenuStr),
			tgbotapi.NewInlineKeyboardButtonData("With quotes", controllers.TArraySecondMenuStr),
		),
	)
	bot.Send(msg)
}

func showTarraySecondMenu(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Choice quotes type:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("'", controllers.TArraySingleQuotes),
			tgbotapi.NewInlineKeyboardButtonData("\"", controllers.TArrayDoubleQuotes),
		),
	)
	bot.Send(msg)
}

func showTarrayStyle(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Choice array style:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("()", controllers.TArrayRound),
			tgbotapi.NewInlineKeyboardButtonData("[]", controllers.TArraySquare),
			tgbotapi.NewInlineKeyboardButtonData("{}", controllers.TArrayCurly),
			tgbotapi.NewInlineKeyboardButtonData("array()", controllers.TArrayOld),
		),
	)
	bot.Send(msg)
}

func sendArray(bot *tgbotapi.BotAPI, chatID int64, arrayType string) {
	tarrayStr := controllers.GetArray(textToArray, quotesType, arrayType)

	msg := tgbotapi.NewMessage(chatID, tarrayStr)
	msg.ParseMode = "HTML"
	bot.Send(msg)
}
