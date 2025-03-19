package menu

import (
	"Qtools/app/controllers"
	"Qtools/app/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func showTtimeText(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "#️⃣ "+helpers.Bold("ttime")+" - a tool for quickly converting timestamp to other time formats \n\n"+
		"Enter timestamp:")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("< back"),
		),
	)
	bot.Send(msg)
}

func handleTtimeMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {
	switch text {
	//case controllers.TArrayFristMenuStr:
	//	quotesType = controllers.TArrayWithoutQuotes
	//	showTarrayStyle(bot, chatID)
	//case controllers.TArraySecondMenuStr:
	//	showTarraySecondMenu(bot, chatID)
	//case controllers.TArraySingleQuotes, controllers.TArrayDoubleQuotes:
	//	quotesType = text
	//	showTarrayStyle(bot, chatID)
	//case controllers.TArrayRound, controllers.TArraySquare, controllers.TArrayCurly, controllers.TArrayOld:
	//	sendArray(bot, chatID, text)
	//	state[chatID] = StateMainMenu
	//	showMainMenu(bot, chatID)
	case "< back":
		state[chatID] = StateMainMenu
		showMainMenu(bot, chatID)
	default:
		getData(bot, chatID, text)
	}
}

func getData(bot *tgbotapi.BotAPI, chatID int64, text string) {
	data := controllers.GetData(text)

	msg := tgbotapi.NewMessage(chatID, data)
	msg.ParseMode = "HTML"
	bot.Send(msg)
}
