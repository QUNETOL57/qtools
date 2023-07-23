package menu

import (
	"Qtools/app/controllers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotState int

const (
	StateMainMenu BotState = iota
	StatePswdMenu
	StateTcaseMenu
	StateTarrayMenu
)

func HandleMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {
	switch state[chatID] {
	case StateMainMenu:
		handleMainMenu(state, bot, chatID, text)
	case StatePswdMenu:
		handlePswdMenu(state, bot, chatID, text)
	case StateTcaseMenu:
		handleTcaseMenu(state, bot, chatID, text)
	case StateTarrayMenu:
		handleTarrayMenu(state, bot, chatID, text)
		//default:
		//	handleUnknownState(bot, chatID)
	}
}

func handleMainMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {
	switch text {
	case "/start":
		showMainMenu(bot, chatID)
		state[chatID] = StateMainMenu
	case "/pswd":
		showPswdMenu(bot, chatID)
		state[chatID] = StatePswdMenu
	case "/tcase":
		showTcaseMenu(bot, chatID)
		state[chatID] = StateTcaseMenu
	case "/tarray":
		showTarrayText(bot, chatID)
		state[chatID] = StateTarrayMenu
		//default:
		//	handleUnknownCommand(bot, chatID)
	}
}

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

func handleTcaseMenu(state map[int64]BotState, bot *tgbotapi.BotAPI, chatID int64, text string) {
	switch text {
	case "< back":
		state[chatID] = StateMainMenu
		showMainMenu(bot, chatID)
	default:
		msg := tgbotapi.NewMessage(chatID, "ok")
		bot.Send(msg)
	}
}
