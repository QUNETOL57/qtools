package services

import (
	"Qtools/app/services/menu"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdates(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	state := make(map[int64]map[menu.BotState]interface{})

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			text := update.Message.Text
			fmt.Println("-----------------")
			fmt.Println(state)
			fmt.Println("-----------------")
			menu.HandleMenu(state, bot, chatID, text)
		} else if update.CallbackQuery != nil {
			menu.HandleInlineMenu(state, bot, update)
		}
	}
}
