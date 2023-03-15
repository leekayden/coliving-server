package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func forbiddenErrMsgController(c *gin.Context) { // temporary solution
	c.JSON(http.StatusOK, gin.H{
		"message": "Resource Forbidden!",
	})
}
