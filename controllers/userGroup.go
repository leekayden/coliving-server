package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Id   uint   `json:"id"`
	Test string `json:"test"`
}

func Register(c *gin.Context) {
	fmt.Println("ran register")
	var requestBody RegisterRequest
	if err := c.BindJSON(&requestBody); err != nil {
		panic("on no error")
	}
	c.Next()
}
