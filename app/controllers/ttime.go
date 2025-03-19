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

	// Загружаем необходимую временную зону (например, Москва)
	moscowLoc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		fmt.Println("Ошибка при загрузке временной зоны Москва:", err)
		return "error"
	}

	tsTime := ConvertFloatTimestamp(tsFloat)
	result := "Greenwich Mean Time (UTC+0): " + helpers.Code(tsTime.UTC().Format("02.01.2006 15:04:05")) + "\nEurope/Moscow (UTC+3): " + helpers.Code(tsTime.In(moscowLoc).Format("02.01.2006 15:04:05"))
	return result
}
