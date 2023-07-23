package menu

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func HandleInlineMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := update.CallbackQuery
	chatID := callback.Message.Chat.ID
	msgId := callback.Message.MessageID
	text := callback.Data

	deleteMessage(bot, chatID, msgId)

	HandleMenu(state, bot, chatID, text)
}

func deleteMessage(bot *tgbotapi.BotAPI, chatId int64, msgId int) {
	deleteMsg := tgbotapi.NewDeleteMessage(chatId, msgId)
	_, err := bot.Send(deleteMsg)
	if err != nil {
		log.Println(err)
	}
}
