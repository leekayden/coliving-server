package router

import (
	"coliving-server/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/test", controllers.TestController)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.Register)
	}
}
