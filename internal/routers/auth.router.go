package routers

import (
	"julianmindria/backendCoffee/internal/handlers"
	"julianmindria/backendCoffee/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/login")

	repo := repositories.NewUser(d)
	handler := handlers.NewAuth(repo)

	route.POST("/auth", handler.Login)

}
