package router

import (
	"coliving-server/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/test", controllers.TestController)
	r.GET("/appname", controllers.AppNameController)
	r.GET("/appcurrency", controllers.AppCurrencyController)
	r.GET("/footertext", controllers.FooterContentController)
	r.GET("/footertitle", controllers.FooterTitleController)
	r.GET("/403msg", controllers.ForbiddenErrMsgController)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.Register)
	}
}
