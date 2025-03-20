package menu

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotState int

const (
	StateMainMenu BotState = iota
	StatePswdMenu
	StateTcaseMenu
	StateTarrayMenu
	StateRstringMenu
	StateTtimeMenu
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
	case StateRstringMenu:
		handleRstringMenu(state, bot, chatID, text)
	case StateTtimeMenu:
		handleTtimeMenu(state, bot, chatID, text)
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
	//case "/tcase":
	//	showTcaseMenu(bot, chatID)
	//	state[chatID] = StateTcaseMenu
	case "/tarray":
		showTarrayText(bot, chatID)
		state[chatID] = StateTarrayMenu
	case "/rstring":
		showRstringText(bot, chatID)
		state[chatID] = StateRstringMenu
	case "/ttime":
		showTtimeText(bot, chatID)
		state[chatID] = StateTtimeMenu
		//default:
		//	handleUnknownCommand(bot, chatID)
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
