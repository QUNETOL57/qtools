package menu

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
