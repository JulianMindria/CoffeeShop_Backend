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
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Message interface{}
}

var Config = cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	AllowCredentials: true,
}
