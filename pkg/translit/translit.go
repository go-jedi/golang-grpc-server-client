package translit

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rob-bender/grpc-new/pkg/helpers"
)

type alphabetColumn struct {
	LetterRusLower string
	LetterEngLower string
}

var alphabet []alphabetColumn = []alphabetColumn{
	{
		LetterRusLower: "а",
		LetterEngLower: "a",
	},
	{
		LetterRusLower: "б",
		LetterEngLower: "b",
	},
	{
		LetterRusLower: "в",
		LetterEngLower: "v",
	},
	{
		LetterRusLower: "г",
		LetterEngLower: "g",
	},
	{
		LetterRusLower: "д",
		LetterEngLower: "d",
	},
	{
		LetterRusLower: "е",
		LetterEngLower: "e",
	},
	{
		LetterRusLower: "ё",
		LetterEngLower: "je",
	},
	{
		LetterRusLower: "ж",
		LetterEngLower: "zh",
	},
	{
		LetterRusLower: "з",
		LetterEngLower: "z",
	},
	{
		LetterRusLower: "и",
		LetterEngLower: "i",
	},
	{
		LetterRusLower: "й",
		LetterEngLower: "y",
	},
	{
		LetterRusLower: "к",
		LetterEngLower: "k",
	},
	{
		LetterRusLower: "л",
		LetterEngLower: "l",
	},
	{
		LetterRusLower: "м",
		LetterEngLower: "m",
	},
	{
		LetterRusLower: "н",
		LetterEngLower: "n",
	},
	{
		LetterRusLower: "о",
		LetterEngLower: "o",
	},
	{
		LetterRusLower: "п",
		LetterEngLower: "p",
	},
	{
		LetterRusLower: "р",
		LetterEngLower: "r",
	},
	{
		LetterRusLower: "с",
		LetterEngLower: "s",
	},
	{
		LetterRusLower: "т",
		LetterEngLower: "t",
	},
	{
		LetterRusLower: "у",
		LetterEngLower: "u",
	},
	{
		LetterRusLower: "ф",
		LetterEngLower: "f",
	},
	{
		LetterRusLower: "х",
		LetterEngLower: "kh",
	},
	{
		LetterRusLower: "ц",
		LetterEngLower: "c",
	},
	{
		LetterRusLower: "ч",
		LetterEngLower: "ch",
	},
	{
		LetterRusLower: "ш",
		LetterEngLower: "sh",
	},
	{
		LetterRusLower: "щ",
		LetterEngLower: "jsh",
	},
	{
		LetterRusLower: "ъ",
		LetterEngLower: "hh",
	},
	{
		LetterRusLower: "ы",
		LetterEngLower: "ih",
	},
	{
		LetterRusLower: "ь",
		LetterEngLower: "jh",
	},
	{
		LetterRusLower: "э",
		LetterEngLower: "eh",
	},
	{
		LetterRusLower: "ю",
		LetterEngLower: "ju",
	},
	{
		LetterRusLower: "я",
		LetterEngLower: "ja",
	},
}

func ConvertRuStringToLatin(needString string) string {
	var strNew string = ""

	for _, value := range needString {
		for _, valueTwo := range alphabet {
			if fmt.Sprintf("%c", value) == valueTwo.LetterRusLower {
				strNew = strNew + valueTwo.LetterEngLower
			}
		}
		if fmt.Sprintf("%c", value) == "_" {
			strNew = strNew + "_"
		}
		if helpers.CheckStringIsInt(fmt.Sprintf("%c", value)) {
			strNew = strNew + fmt.Sprintf("%c", value)
		}
	}

	return strNew
}

func ConvertRuStringNameFileToLatin(fileName string) string {
	var extension = filepath.Ext(fileName)
	var name = strings.ToLower(strings.ReplaceAll(regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9 ]+`).ReplaceAllString(fileName[0:len(fileName)-len(extension)], ""), " ", "_"))
	var strNew string = ""

	for _, value := range name {
		for _, valueTwo := range alphabet {
			if fmt.Sprintf("%c", value) == valueTwo.LetterRusLower {
				strNew = strNew + valueTwo.LetterEngLower
			}
		}
		if fmt.Sprintf("%c", value) == "_" || fmt.Sprintf("%c", value) == "." {
			strNew = strNew + "_"
		}
		if helpers.CheckStringIsInt(fmt.Sprintf("%c", value)) {
			strNew = strNew + fmt.Sprintf("%c", value)
		}
		matched, _ := regexp.MatchString("[a-zA-Z]", fmt.Sprintf("%c", value))
		if matched {
			strNew = strNew + strings.ToLower(fmt.Sprintf("%c", value))
		}
	}

	return fmt.Sprintf("%s%s", strNew, extension)
}
