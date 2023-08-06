package menu

import (
	"Qtools/app/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleDsymbMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {

	switch text {
	case "< back":
		state[chatID] = StateMainMenu
		showMainMenu(bot, chatID)
	default:
		textToArray = text
		showTarrayFirstMenu(bot, chatID)
	}
}

//func handleDsymbMenuSteps(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {
//	switch state[chatID][] {
//	case :
//	case "< back":
//		state[chatID] = StateMainMenu
//		showMainMenu(bot, chatID)
//	default:
//		textToArray = text
//		showTarrayFirstMenu(bot, chatID)
//	}
//}

func showDsymbText(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "#️⃣ "+helpers.Bold("dsymb")+" - dsymb is a tool for quickly removing all occurrences of a specific character, word, or sentence from a given text \n\n"+
		"Enter text:")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("< back"),
		),
	)
	bot.Send(msg)
}
