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
	route.PUT("/:username", middleware.UploadFile, middleware.AuthJwt("admin"), handler.Updatedata)
	route.DELETE("/", middleware.AuthJwt("admin"), handler.Deletedata)
}
