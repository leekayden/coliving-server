package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCorsConfig() gin.HandlerFunc {
	frontend := os.Getenv("frontend")
	return cors.New(cors.Config{
		AllowOrigins:     []string{frontend},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			fmt.Println(origin)
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	})
}