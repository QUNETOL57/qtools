package controllers

import (
	"Qtools/app/helpers"
	"fmt"
	"math"
	"strconv"
	"time"
)

func ConvertFloatTimestamp(ts float64) time.Time {
	// Отделяем целую часть (секунды)
	secs := int64(ts)
	// Вычисляем дробную часть и переводим её в наносекунды
	nsecs := int64(math.Round((ts - float64(secs)) * 1e9))
	return time.Unix(secs, nsecs)
}

func GetData(ts string) string {
	tsFloat, err := strconv.ParseFloat(ts, 64)
	if err != nil {
		return "error"
	}

	moscowLoc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		fmt.Println("Ошибка при загрузке временной зоны Москва:", err)
		return "error"
	}

	tsTime := ConvertFloatTimestamp(tsFloat)
	result := "Greenwich Mean Time (UTC+0)" +
		"\n========================================" +
		"\n" + helpers.Code("RU DateTime | ") + helpers.Code(tsTime.UTC().Format("02.01.2006 15:04:05")) +
		"\n" + helpers.Code("DateTime    | ") + helpers.Code(tsTime.UTC().Format(time.DateTime)) +
		"\n" + helpers.Code("RFC3339     | ") + helpers.Code(tsTime.UTC().Format(time.RFC3339)) +
		"\n" + helpers.Code("RFC1123     | ") + helpers.Code(tsTime.UTC().Format(time.RFC1123)) +
		"\n" + helpers.Code("ANSIC       | ") + helpers.Code(tsTime.UTC().Format(time.ANSIC)) +
		"\n\nEurope/Moscow (UTC+3):" +
		"\n========================================" +
		"\n" + helpers.Code("RU DateTime | ") + helpers.Code(tsTime.In(moscowLoc).Format("02.01.2006 15:04:05")) +
		"\n" + helpers.Code("DateTime    | ") + helpers.Code(tsTime.In(moscowLoc).Format(time.DateTime)) +
		"\n" + helpers.Code("RFC3339     | ") + helpers.Code(tsTime.In(moscowLoc).Format(time.RFC3339)) +
		"\n" + helpers.Code("RFC1123     | ") + helpers.Code(tsTime.In(moscowLoc).Format(time.RFC1123)) +
		"\n" + helpers.Code("ANSIC       | ") + helpers.Code(tsTime.In(moscowLoc).Format(time.ANSIC))
	return result
}
