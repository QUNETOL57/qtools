package helpers

import "strings"

func StringToArray(text string) ([]string, string) {
	textArray := strings.Split(strings.TrimSpace(text), " ")
	separator := " "

	if len(textArray) == 1 {
		textArray = strings.Split(text, "\n")
		separator = "\n"
	}
	return textArray, separator
}
