package router

import (
	"coliving-server/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/test", controllers.testController)
	r.GET("/appname", controllers.appNameController)
	r.GET("/appcurrency", controllers.appCurrencyController)
	r.GET("/footertext", controllers.footerTextController)
	r.GET("/footertitle", controllers.footerTitleController)
	r.GET("/403msg", controllers.forbiddenErrMsgController)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.Register)
	}
}
