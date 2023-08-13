package config

import (
	"context"

	"github.com/gin-contrib/cors"
)

type Operation func(ctx context.Context) error

type Metas struct {
	Next  interface{} `json:"next"`
	Prev  interface{} `json:"prev"`
	Total int         `json:"total"`
}

type Result struct {
	Data    interface{}
	Meta    interface{}
	Message interface{}
}

var CorsConfig = cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	AllowCredentials: true,
}
