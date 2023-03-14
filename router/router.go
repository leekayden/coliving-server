package router

import (
	"coliving-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	controllers.Test(router)
}
