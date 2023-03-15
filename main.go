package main

import (
	"coliving-server/config"
	"coliving-server/database"
	leRouter "coliving-server/router"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load(".env")
	database.InitDatabase()

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	router.Use(config.NewCorsConfig())

	leRouter.SetUpRoutes(router)
	router.Run(":3003")
	fmt.Println("Started")
}
