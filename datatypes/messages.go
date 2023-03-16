package datatypes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// errors: >= 1000
const (
	// general (1000)
	ErrorCodeGeneric             = 1000
	ErrorCodeIncorrectParameters = 1001

	// user related (2000)
	ErrorCodeEmailExists = 2000
)

const (
	ServerErrorGeneric             = "Something went wrong with the server, please try again later."
	ServerErrorIncorrectParameters = "The server failed to get the parameters."

	ServerErrorEmailExists = "The email you are using to register already exists in the database. Do you want to login instead?"
)

var errMsgTable = map[int]string{
	ErrorCodeGeneric:             ServerErrorGeneric,
	ErrorCodeIncorrectParameters: ServerErrorIncorrectParameters,
	ErrorCodeEmailExists:         ServerErrorEmailExists,
}

type Msi map[string]interface{}

type ResponseError struct {
	Code    uint   `json:"code"` // will be our own, refer to above
	Message string `json:"message"`
}

func GenericError(c *gin.Context) {
	RespondWithError(c, ErrorCodeGeneric)
}

func RespondWithError(c *gin.Context, internalCode int) {
	c.JSON(http.StatusInternalServerError, ResponseError{Code: uint(internalCode), Message: errMsgTable[internalCode]})
}

type ResponseSuccess struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	SuccessCodeGeneric = 1

	SuccessCodeRegistered = 10
)

const (
	SuccessMsgGeneric    = "Server responded successfully."
	SuccessMsgRegistered = "Successfully registered user."
)

var successMsgTable = map[int]string{
	SuccessCodeGeneric:    SuccessMsgGeneric,
	SuccessCodeRegistered: SuccessMsgRegistered,
}

func RespondWithSuccess(c *gin.Context, internalCode int, data interface{}) {
	if internalCode >= 1000 {
		log.Warn().
			Str("test", "RespondWithSuccess should receive internalCode of below 1000, instead received "+strconv.FormatInt(int64(internalCode), 10))
	}
	c.JSON(http.StatusOK, ResponseSuccess{Code: uint(internalCode), Message: successMsgTable[internalCode], Data: data})
}
