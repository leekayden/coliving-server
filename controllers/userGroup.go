package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	fmt.Println("ran register")
	c.Next()
}
