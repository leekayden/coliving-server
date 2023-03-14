package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(router *gin.Engine) { // temporary solution
	fmt.Println("what")
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
}
