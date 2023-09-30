package controllers

import (
	"Qtools/app/helpers"
)

func RemoveDuplicate(text string) (string, string) {
	seen := make(map[string]bool)
	duplicatesArr := []string{}

	textArr, separator := helpers.StringToArray(text)

	for i, word := range textArr {
		if seen[word] {
			duplicatesArr = append(duplicatesArr, word)
			textArr[i] = ""
		} else {
			seen[word] = true
		}
	}

	newTextStr := genTextStr(textArr, separator)
	duplicatesStr := genTextStr(duplicatesArr, " ")

	if len(textArr) > 2 {
		newTextStr = newTextStr[:len(newTextStr)-1]
	}
	if len(duplicatesArr) > 2 {
		duplicatesStr = duplicatesStr[:len(newTextStr)-1]
	}

	return newTextStr, duplicatesStr
}

func genTextStr(wordArr []string, separator string) string {
	textStr := ""
	for _, word := range wordArr {
		if word == "" {
			continue
		}
		textStr += word + separator
	}

	return textStr
}
