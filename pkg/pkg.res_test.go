package pkg

import (
	"julianmindria/backendCoffee/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	code     = 500
	codefail = 2023
	desc     = ""
)

func TestGetStatus(t *testing.T) {
	t.Run("cek success", func(t *testing.T) {
		var response = getStatus(code)
		desc = response
		assert.NotEqual(t, "", response, "Nil")
	})

	t.Run("cek failed", func(t *testing.T) {
		var response = getStatus(codefail)
		assert.Equal(t, "", response, "Description is not nil")
	})
}

func TestNewRes(t *testing.T) {
	t.Run("cek", func(t *testing.T) {
		expected := Response{Code: code, Status: desc}
		var res = NewRes(code, &config.Result{})
		assert.Equal(t, expected, *res, "Response is not correct")
	})
}
