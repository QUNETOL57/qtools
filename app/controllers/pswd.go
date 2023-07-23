package controllers

import (
	"Qtools/app/helpers"
	"math/rand"
	"strconv"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var lettersExtra = append(letters, []rune("!@#$%^&*")...)

type PasswordsOptions struct {
	CountChar      int
	CountPasswords int
	WithSymbols    bool
}

func BasePasswordOptions() PasswordsOptions {
	return PasswordsOptions{
		CountChar:      8,
		CountPasswords: 5,
		WithSymbols:    false,
	}
}

func genPasswords(options PasswordsOptions) []string {
	var pswdArr []string

	for i := 1; i <= options.CountPasswords; i++ {
		password := make([]rune, options.CountChar)

		for c := range password {
			if options.WithSymbols {
				password[c] = lettersExtra[rand.Intn(len(letters))]
			} else {
				password[c] = letters[rand.Intn(len(letters))]
			}
		}

		pswdArr = append(pswdArr, string(password))
	}

	return pswdArr
}

func GetPswd() string {
	pswdArr := genPasswords(BasePasswordOptions())
	pswdHtml := "Список сгенерированных паролей:\n"

	for i, password := range pswdArr {
		pswdHtml += strconv.Itoa(i+1) + "." + helpers.Code(string(password)) + "\n"
	}

	return pswdHtml
}
