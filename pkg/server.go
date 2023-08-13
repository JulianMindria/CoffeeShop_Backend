package pkg

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Server(router *gin.Engine) *http.Server {
	var addr string = "0.0.0.0:8080"
	if port := viper.GetString("PORT"); port != "" {
		addr = ":" + port
	}

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	return srv
}
