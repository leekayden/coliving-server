package main

import (
	"coliving-server/config"
	leRouter "coliving-server/router"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load(".env")
	router := gin.Default()
	router.Use(config.NewCorsConfig())
	leRouter.SetUpRoutes(router)
	router.Run(":3003") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Started")
}
