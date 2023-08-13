package routers

import (
	"julianmindria/backendCoffee/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(config.CorsConfig))

	User(router, db)
	Product(router, db)
	auth(router, db)

	return router
}
