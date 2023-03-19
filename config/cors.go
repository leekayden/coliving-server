package config

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{Globals.Frontend},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// fmt.Println(origin)
			// return origin == "https://github.com"
			return true;
		},
		MaxAge: 12 * time.Hour,
	})
}
