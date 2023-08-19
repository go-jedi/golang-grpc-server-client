package translit

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConvertRuStringToLatin(t *testing.T) {
	// Arrange
	var testTable []struct {
		name       string
		needString string
		expected   string
	} = []struct {
		name       string
		needString string
		expected   string
	}{
		{
			name:       "OK",
			needString: strings.ToLower("Делаю тестирование"),
			expected:   "delajutestirovanie",
		},
		{
			name:       "OK (BROADCAST LATIN)",
			needString: strings.ToLower("doing testing"),
			expected:   "",
		},
		{
			name:       "Error (NOT LOWER CASE)",
			needString: "Делаю тестирование",
			expected:   "elajutestirovanie",
		},
	}
	// Act
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := ConvertRuStringToLatin(test.needString)
			// Assert
			assert.Equal(t, test.expected, result, fmt.Sprintf("Incorrect result. Expect %s, got %s",
				test.expected,
				result,
			))
		})
	}
}

func Test_ConvertRuStringNameFileToLatin(t *testing.T) {
	// Arrange
	var testTable []struct {
		name     string
		fileName string
		expected string
	} = []struct {
		name     string
		fileName string
		expected string
	}{
		{
			name:     "OK",
			fileName: "Делаю-тестирование.txt",
			expected: "delajutestirovanie.txt",
		},
		{
			name:     "OK (BROADCAST LATIN)",
			fileName: "doing-testing.txt",
			expected: "doingtesting.txt",
		},
		{
			name:     "OK (BROADCAST EMPTY LINE)",
			fileName: "",
			expected: "",
		},
	}
	// Act
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := ConvertRuStringNameFileToLatin(test.fileName)
			// Assert
			assert.Equal(t, test.expected, result, fmt.Sprintf("Incorrect result. Expect %s, got %s",
				test.expected,
				result,
			))
		})
	}
}
