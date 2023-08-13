package routers

import (
	"julianmindria/backendCoffee/internal/handlers"
	"julianmindria/backendCoffee/internal/middleware"
	"julianmindria/backendCoffee/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func User(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", middleware.UploadFile, handler.Postdata)
	route.GET("/", middleware.AuthJwt("user"), handler.Getdata)
	route.PUT("/", handler.Updatedata)
	route.DELETE("/", handler.Deletedata)
}
