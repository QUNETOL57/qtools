package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
	"strings"
)

var numericKeyboardTCase = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("CamelCase", "tcase1"),
		tgbotapi.NewInlineKeyboardButtonData("snake_case", "tcase2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("SCREAMING_SNAKE_CASE", "tcase3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("PascalCase", "tcase4"),
		tgbotapi.NewInlineKeyboardButtonData("kebab-case", "tcase5"),
	),
)

func convertSnakeToCamel(snakeCase string) string {
	words := strings.Split(snakeCase, "_")
	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func checkCase(text string) string {
	switch {
	case checkSnakeCase(text):
		return "snake"
	case checkKebabCase(text):
		return "kebab"
	case checkPascalCase(text):
		return "pascal"
	case checkCamelCase(text):
		return "camel"
	case checkScreamingSnakeCase(text):
		return "scream"
	default:
		return "plane"
	}
}

func SetTCaseSettings(update tgbotapi.Update) tgbotapi.MessageConfig {
	//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите тип генерации:")
	//msg.ReplyMarkup = numericKeyboardTCase

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, checkCase(update.Message.Text))

	return msg
}

func checkCamelCase(text string) bool {
	regex := regexp.MustCompile("^[a-zA-Z]+(?:[A-Z][a-z]*)*$")
	return regex.MatchString(text)
}

func checkPascalCase(text string) bool {
	regex := regexp.MustCompile("^[A-Z][a-zA-Z]*$")
	return regex.MatchString(text)
}

func checkScreamingSnakeCase(text string) bool {
	regex := regexp.MustCompile("^[A-Z0-9_]+$")
	return regex.MatchString(text)
}

func checkSnakeCase(text string) bool {
	regex := regexp.MustCompile("^[a-z0-9_]+$")
	return regex.MatchString(text)
}

func checkKebabCase(text string) bool {
	regex := regexp.MustCompile("^[a-z0-9]+(-[a-z0-9]+)*$")
	return regex.MatchString(text)
}
