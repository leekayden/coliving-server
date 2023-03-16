package main

import (
	"coliving-server/config"
	"coliving-server/database"
	"coliving-server/datatypes"
	leRouter "coliving-server/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	database.InitDatabase()
	config.InitGlobals()

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
		// if err, ok := recovered.(string); ok {
		// 	c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		// }
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": datatypes.ServerErrorGeneric,
		})
	}))

	router.Use(config.NewCorsConfig())

	leRouter.SetUpRoutes(router)
	router.Run(":3003")
	fmt.Println("Started")
}
