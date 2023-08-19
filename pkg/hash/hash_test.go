package hash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GeneratePasswordHash(t *testing.T) {
	// Arrange
	var testTable []struct {
		name     string
		password string
		expected string
	} = []struct {
		name     string
		password string
		expected string
	}{
		{
			name:     "OK",
			password: "test",
			expected: "686a7172686a7177313234363137616a6668616a739f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		},
		{
			name:     "ERROR (NO EQUAL)",
			password: "test",
			expected: "686a7172686a7177313234363137616a6668616a739f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f121312",
		},
	}
	// Act
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result, _ := GeneratePasswordHash(test.password)
			// Assert
			switch test.name {
			case "OK":
				assert.Equal(t, test.expected, result, fmt.Sprintf("Incorrect result. Expect %s, got %s",
					test.expected,
					result,
				))
			case "ERROR (NO EQUAL)":
				assert.NotEqual(t, test.expected, result)
			}
		})
	}
}
