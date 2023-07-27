package controllers

import (
	"Qtools/app/helpers"
	"strings"
)

type TArrayMenuState int

const (
	TArrayFristMenu TArrayMenuState = iota
	TArraySecondMenu
)

const (
	TArrayFristMenuStr  = "tarray_first"
	TArraySecondMenuStr = "tarray_second"
	TArraySingleQuotes  = "tarray_single_quotes"
	TArrayDoubleQuotes  = "tarray_double_quotes"
	TArrayWithoutQuotes = "tarray_without_quotes"
	TArrayRound         = "tarray_round"
	TArraySquare        = "tarray_square"
	TArrayCurly         = "tarray_curly"
	TArrayOld           = "tarray_old"
)

func setArray(text string, quotesType string, bracketsType string) string {
	var quote string
	var brackets string
	var textArrayStr string
	textArray := strings.Split(strings.TrimSpace(text), " ")

	if len(textArray) == 1 {
		textArray = strings.Split(text, "\n")
	}

	switch quotesType {
	case TArraySingleQuotes:
		quote = "'"
	case TArrayDoubleQuotes:
		quote = "\""
	case TArrayWithoutQuotes:
		quote = ""
	}

	switch bracketsType {
	case TArrayRound:
		brackets = "({text})"
	case TArraySquare:
		brackets = "[{text}]"
	case TArrayCurly:
		brackets = "{{text}}"
	case TArrayOld:
		brackets = "array({text})"
	}

	for _, value := range textArray {
		textArrayStr += quote + value + quote + ", "
	}
	return strings.ReplaceAll(brackets, "{text}", textArrayStr[:len(textArrayStr)-2])
}

func GetArray(text string, quotesType string, arrayType string) string {
	tarrayStr := setArray(text, quotesType, arrayType)

	return "Array : \n" + helpers.Code(tarrayStr)
}
