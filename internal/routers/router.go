package routers

import (
	"julianmindria/backendCoffee/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRoute(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())
	User(router, db)
	Product(router, db)
	auth(router, db)

	return router
}
