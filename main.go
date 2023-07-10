package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
	rand.Seed(time.Now().UnixNano())
}

func genPasswords(countChar int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	var pswdHtml string = ""

	for i := 1; i <= 5; i++ {
		password := make([]rune, countChar)

		for c := range password {
			password[c] = letters[rand.Intn(len(letters))]
		}

		pswdHtml += strconv.Itoa(i) + ". <code>" + string(password) + "</code>\n"
	}

	return pswdHtml
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
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		switch update.Message.Command() {
		case "pswd":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, genPasswords(10))
			msg.ParseMode = "HTML"
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
