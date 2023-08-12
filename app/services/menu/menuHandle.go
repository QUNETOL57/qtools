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
)


type Menu struct {
    Items     map[BotState]MenuItem // Карта элементов меню, ключ - идентификатор элемента
}

type MenuItem struct {
    Text         string            // Текст элемента меню
    NextMenuId   string            // Идентификатор следующего подменю
    Command      func()            // Функция-обработчик команды, связанной с элементом меню
	Handle 		 func()
	SubItems	 map[BotState]MenuItem

}

func (m *Menu) GetUserState(userId string) string {
    return m.UserState[userId]
}

func (m *Menu) HandleCommand(userId string, command string) {
    currentItem := m.Items[m.GetUserState(userId)] // Получаем текущий элемент меню

    if currentItem.Command != nil {
        currentItem.Command() // Выполняем функцию-обработчик команды текущего элемента меню
    }

    nextMenuId := currentItem.NextMenuId // Получаем идентификатор следующего подменю

    if nextMenuId != "" {
        m.SetUserState(userId, nextMenuId) // Устанавливаем состояние пользователя как идентификатор следующего подменю
    }
}

func main() {
    menu := Menu{
        Items:     make(map[BotState]MenuItem){
			StateMainMenu : MenuItem{
				Text: "Main menu",
				Command: func() { fmt.Println("Вывод главного меню") }
				Handle: func() { fmt.Println("Обработчик для меню") }
			},
			StatePswdMenu : MenuItem{
				Text: "pwsd",
				Command: func() { fmt.Println("Вывод меню pswd") }
				SubItems: map[BotState]MenuItem{
					StatePswdMenu1 : MenuItem{
						Text: "pwsd simple",
						Command: func() { fmt.Println("Вывод simple") }
						NextMenuId: StateMainMenu
					},
					StatePswdMenu2 : MenuItem{
						Text: "pwsd hard",
						Command: func() { fmt.Println("Вывод hard") }
						SubItems: map[BotState]MenuItem{
							StatePswdMenu3 : MenuItem{
								Text: "pwsd настроить количество символов",
								Command: func() { fmt.Println("Вывод simple") }
								NextMenuId: StatePswdMenu2
							},
							StatePswdMenu4 : MenuItem{
								Text: "pwsd настроить количество паролей",
								Command: func() { fmt.Println("Вывод simple") }
								NextMenuId: StatePswdMenu2
							},
							StatePswdMenu5 : MenuItem{
								Text: "pwsd настроить вывод спец символов",
								Command: func() { fmt.Println("Вывод simple") }
								NextMenuId: StatePswdMenu2
							},
						}
					},
			},
			
		}
    }
	
	userState := [chatID: StatePswdMenu3]
    // Добавляем элементы меню
    menu.AddItem("main", MenuItem{Text: "Main menu",  Command: func() { fmt.Println("") }})
    menu.AddItem("sub", MenuItem{Text: "Подменю", Command: func() { fmt.Println("Выполняется команда из подменю") }})

    // Устанавливаем начальное состояние пользователя
    menu.SetUserState("user1", "main")

    // Обрабатываем команду пользователя
    menu.HandleCommand("user1", "")

    // Вывод результата в консоль: "Выполняется команда из подменю"
}

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
	//case "/tcase":
	//	showTcaseMenu(bot, chatID)
	//	state[chatID] = StateTcaseMenu
	case "/tarray":
		showTarrayText(bot, chatID)
		state[chatID] = StateTarrayMenu
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
