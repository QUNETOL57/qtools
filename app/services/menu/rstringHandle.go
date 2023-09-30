package menu

import (
	"Qtools/app/controllers"
	"Qtools/app/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleRstringMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {
	textNew, textOld, duplicates := controllers.RemoveDuplicate(text)

	msg := tgbotapi.NewMessage(chatID, "Text without duplicates:\n"+textNew)
	bot.Send(msg)
	msg2 := tgbotapi.NewMessage(chatID, "Duplicates:\n"+helpers.Code(duplicates)+"\n\nOriginal text with highlighted duplicates:\n"+textOld)
	msg2.ParseMode = "HTML"
	bot.Send(msg2)

	state[chatID] = StateMainMenu
	showMainMenu(bot, chatID)
}

func showRstringText(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "#️⃣ "+helpers.Bold("rstring")+" - a tool for removing duplicates in text \n\n"+
		"Please enter text where words are separated by a space or a line break :")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("< back"),
		),
	)
	bot.Send(msg)
}
