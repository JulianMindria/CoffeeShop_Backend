package middleware

import (
	"julianmindria/backendCoffee/config"
	"julianmindria/backendCoffee/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwt(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			pkg.NewRes(401, &config.Result{
				Message: "Please login",
			}).Send(ctx)
			return
		}

		if !strings.Contains(header, "Bearer") {
			pkg.NewRes(401, &config.Result{
				Message: "Header Invalid",
			}).Send(ctx)
			return
		}

		tokens := strings.Replace(header, "Bearer ", "", -1)

		check, err := pkg.Verifytoken(tokens)
		if err != nil {
			pkg.NewRes(401, &config.Result{
				Message: err.Error(),
			}).Send(ctx)
			return
		}

		for _, r := range role {
			if r == check.Role {
				valid = true
			}
		}

		if !valid {
			pkg.NewRes(401, &config.Result{
				Message: "Not authorized",
			}).Send(ctx)
			return
		}
		ctx.Set("userID", check.Id)
		ctx.Next()
	}
}
