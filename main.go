package main

import (
	leRouter "coliving-server/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	leRouter.SetUpRoutes(router)
	router.Run(":3003") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Started")
}
