package pkg

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var token = ""
var uuid = "1122334455667788"
var role = "user"
var data_token = claims{}

func Test_newToken(t *testing.T) {
	t.Run("Set token", func(t *testing.T) {
		var expected = claims{Id: uuid, Role: role, RegisteredClaims: jwt.RegisteredClaims{Issuer: "coffeeShop"}}
		var response = NewToken(uuid, role)
		response.RegisteredClaims.ExpiresAt = nil
		data_token = *response
		assert.Equal(t, expected, *response, "Data token results do not match")
	})
}

func Test_GenToken(t *testing.T) {
	t.Run("generate token", func(t *testing.T) {
		response, err := data_token.Generate()
		token = response
		assert.NotEqual(t, "", response, "Empty token")
		assert.Nil(t, err, "Generating token error")
	})
}

func Test_Verify(t *testing.T) {
	t.Run("verify token success", func(t *testing.T) {
		var expected = data_token
		response, err := Verifytoken(token)
		assert.Equal(t, expected, *response, "Data token results do not match")
		assert.Nil(t, err, "Generating token error")
	})
}
