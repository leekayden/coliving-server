package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	id uint `json:"id"`
}

func Register(c *gin.Context) {
	fmt.Println("ran register")
	var requestBody RegisterRequest
	if err := c.BindJSON(&requestBody); err != nil {
		panic("on no error")
	}
	fmt.Printf("%#v\n", requestBody.id)
	c.Next()
}
