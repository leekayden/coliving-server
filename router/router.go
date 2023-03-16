package router

import (
	"coliving-server/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://edenco.vercel.app/404")
	})
	r.GET("/test", controllers.TestController)
	r.GET("/appname", controllers.AppNameController)
	r.GET("/appcurrency", controllers.AppCurrencyController)
	r.GET("/footercontent", controllers.FooterContentController)
	r.GET("/footertitle", controllers.FooterTitleController)
	r.GET("/403msg", controllers.ForbiddenErrMsgController)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.Register)
	}
}
