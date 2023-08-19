package bcrypt_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rob-bender/grpc-new/pkg/bcrypt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_HashPassword(t *testing.T) {
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
			expected: "$2a$14$ux1dOZJriumbEIheHryf1eA6qwa0qG2j3bnjU0i2g3KUbD5QqEUr.",
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
			result, err := bcrypt.HashPassword(test.password)
			// Assert
			switch test.name {
			case "OK":
				require.NoError(t, err)
				assert.Equal(t, strings.Contains(test.expected, "$2a$14$"), strings.Contains(result, "$2a$14$"), fmt.Sprintf("Incorrect result. Expect %s, got %s",
					test.expected,
					result,
				))
			case "ERROR (NO EQUAL)":
				require.NoError(t, err)
				assert.NotEqual(t, test.expected, result)
			}
		})
	}
}
