package menu

import (
	"Qtools/app/controllers"
	"Qtools/app/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleRstringMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {
	newText, dublicates := controllers.RemoveDuplicate(text)

	msg := tgbotapi.NewMessage(chatID, newText)
	bot.Send(msg)
	msg2 := tgbotapi.NewMessage(chatID, dublicates)
	bot.Send(msg2)

	state[chatID] = StateMainMenu
	showMainMenu(bot, chatID)
}

func showRstringText(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "#️⃣ "+helpers.Bold("rstring")+" - a tool for removing duplicates in text \n\n"+
		"Enter text :")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("< back"),
		),
	)
	bot.Send(msg)
}
