package routers

import (
	"julianmindria/backendCoffee/internal/handlers"
	"julianmindria/backendCoffee/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.NewProduct(d)
	handler := handlers.NewProduct(repo)
	otherRepo := repositories.NewFav(d)
	otherhandler := handlers.NewFav(otherRepo)

	route.POST("/", handler.Postdata)
	route.POST("/favorite", otherhandler.Addfavorite)
	route.GET("/", handler.Getdata)
	route.PUT("/", handler.Updatedata)
	route.DELETE("/", handler.Deletedata)
}
