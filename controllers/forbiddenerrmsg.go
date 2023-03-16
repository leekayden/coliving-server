package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ForbiddenErrMsgController(c *gin.Context) { // temporary solution
	c.JSON(http.StatusOK, gin.H{
		"message": "Resource Forbidden!",
	})
}
