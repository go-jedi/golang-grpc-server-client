package helpers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func getCurrentDate() time.Time { // получить текущее время
	return time.Now()
}

func ContainsInt(a []int, x int) bool { // проверить существует ли в массиве чисел нужное нам число
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ContainsString(a []string, x string) bool { // проверить существует ли в массиве строк нужная нам строка
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func UnicodeToutf8(source string) string {
	var res = []string{""}
	sUnicode := strings.Split(source, "\\u")
	var context = ""
	for _, v := range sUnicode {
		var additional = ""
		if len(v) < 1 {
			continue
		}
		if len(v) > 4 {
			rs := []rune(v)
			v = string(rs[:4])
			additional = string(rs[4:])
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			context += v
		}
		context += fmt.Sprintf("%c", temp)
		context += additional
	}
	res = append(res, context)
	return strings.Join(res, "")
}

func CheckExistDir(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) { // проверка существует ли папка или нет
		return true
	} else {
		return false
	}
}

func CheckStringIsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
