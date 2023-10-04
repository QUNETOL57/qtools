package controllers

import (
	"Qtools/app/helpers"
)

func RemoveDuplicate(text string) (string, string, string) {
	seen := make(map[string]bool)
	duplicatesArr := []string{}

	textArrNew, separator := helpers.StringToArray(text)
	textArrOld := make([]string, len(textArrNew))

	for i, word := range textArrNew {
		if seen[word] {
			duplicatesArr = append(duplicatesArr, word)
			textArrNew[i] = ""
			textArrOld[i] = helpers.Code(word)
		} else {
			seen[word] = true
			textArrOld[i] = word
		}
	}

	textStrNew := genTextStr(textArrNew, separator)
	textStrOld := genTextStr(textArrOld, separator)
	duplicatesStr := genTextStr(duplicatesArr, " ")

	if len(textArrNew) > 2 {
		textStrNew = textStrNew[:len(textStrNew)-1]
	}
	if len(duplicatesArr) > 2 {
		duplicatesStr = duplicatesStr[:len(duplicatesStr)-1]
	}
	if len(textArrOld) > 2 {
		textStrOld = textStrOld[:len(textStrOld)-1]
	}

	return textStrNew, textStrOld, duplicatesStr
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
