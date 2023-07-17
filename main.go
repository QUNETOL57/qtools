package main

import (
	"Qtools/app/controllers"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	telegramToken, exists := os.LookupEnv("TELEGRAM_TOKEN")

	if exists {
		fmt.Println(telegramToken)
	}

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "/pswd":
				msg = controllers.SetPswdSettings(update)
			case "/tcase":
				msg = controllers.SetTCaseSettings(update)
			}

			msg = controllers.SetTCaseSettings(update)

			sendMsg, err := bot.Send(msg)
			if err != nil {
				panic(err)
			}

			messageID := sendMsg.MessageID

			fmt.Println(messageID)
		} else if update.CallbackQuery != nil {
			callback := update.CallbackQuery

			deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)
			_, err := bot.Send(deleteMsg)
			if err != nil {
				log.Println(err)
			}

			//callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "hello")
			//if _, err := bot.Request(callback); err != nil {
			//	panic(err)
			//}
			msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "")
			switch callback.Data {
			case "pswd1":
				msg.Text = controllers.GetPswd()
				msg.ParseMode = "HTML"
			}

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}

		}
	}
}
