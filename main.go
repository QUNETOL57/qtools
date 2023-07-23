package main

import (
	"Qtools/app/services"

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

	go services.HandleUpdates(bot)

	select {}
}
