package controllers

import (
	"coliving-server/database/handlers"
	"coliving-server/database/models"
	"coliving-server/datatypes"
	"coliving-server/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	FirstName        string `form:"first_name"`
	LastName         string `form:"last_name"`
	Location         string `form:"location"`
	Email            string `form:"email"`
	Password         string `form:"password"`
	ReceiveAddEmails bool   `form:"receive_add_emails"`
}

func Register(c *gin.Context) {
	fmt.Println("ran register")
	var requestBody RegisterRequest
	if err := c.BindJSON(&requestBody); err != nil {
		datatypes.GenericError(c)
		return
	}
	hashedPassword, err := services.HashAndSaltPassword(requestBody.Password)
	if err != nil {
		datatypes.GenericError(c)
		return
	}
	exists, err := handlers.DoesEmailExist(requestBody.Email)
	if err != nil {
		datatypes.GenericError(c)
		return
	}
	if exists {
		datatypes.RespondWithError(c, datatypes.ErrorCodeEmailExists)
		return
	}
	handlers.CreateUser(&models.User{
		FirstName:        requestBody.FirstName,
		LastName:         requestBody.LastName,
		Location:         requestBody.Location,
		Email:            requestBody.Email,
		Password:         hashedPassword,
		ReceiveAddEmails: requestBody.ReceiveAddEmails,
	})
	datatypes.RespondWithSuccess(c, datatypes.SuccessCodeRegistered, nil)
	c.Next()
}
