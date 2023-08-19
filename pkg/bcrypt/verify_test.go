package bcrypt_test

import (
	"testing"

	"github.com/rob-bender/grpc-new/pkg/bcrypt"
	"github.com/stretchr/testify/assert"
)

func Test_VerifyPassword(t *testing.T) {
	// Arrange
	var testTable []struct {
		name             string
		userPassword     string
		providedPassword string
		expected         bool
	} = []struct {
		name             string
		userPassword     string
		providedPassword string
		expected         bool
	}{
		{
			name:             "OK",
			userPassword:     "test",
			providedPassword: "$2a$14$ux1dOZJriumbEIheHryf1eA6qwa0qG2j3bnjU0i2g3KUbD5QqEUr.",
			expected:         true,
		},
		{
			name:             "ERROR (NO EQUAL)",
			userPassword:     "test",
			providedPassword: "$2a$14$ux1dOZJriumbEIheHryf1eA6qwa0qG2j3bnjU0i2g3KUbD5Qq",
			expected:         false,
		},
	}
	// Act
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result, _ := bcrypt.VerifyPassword(test.userPassword, test.providedPassword)
			// Assert
			switch test.name {
			case "OK":
				assert.Equal(t, test.expected, result)
			case "ERROR (NO EQUAL)":
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
