package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var password = "12345678"
var hasedPassword string
var errors error

func TestHashPassword(t *testing.T) {
	hasedPassword, errors = HashPassword(password)
	assert.NoError(t, errors, "error occured while hashing password")
	assert.NotEqual(t, password, hasedPassword, "Hashing Error")
}

func TestVerifyPassword(t *testing.T) {
	t.Run("verify success", func(t *testing.T) {
		var hasPassword = VerifyPassword(hasedPassword, password)
		assert.Nil(t, hasPassword, "Wrong Password")
	})

	t.Run("verify failed", func(t *testing.T) {
		var hasPassword = VerifyPassword(hasedPassword, "julianmindria")
		assert.NotNil(t, hasPassword, "Correct Password")
	})
}
