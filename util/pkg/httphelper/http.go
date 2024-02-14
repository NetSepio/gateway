package httphelper

import (
	"net/http"

	"github.com/NetSepio/gateway/api/types"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

//TODO: add method for internet error with common msg
func ErrResponse(c *gin.Context, statusCode int, errMessage string) {
	response := types.ApiResponse{
		StatusCode: statusCode,
		Error:      errMessage,
	}
	c.JSON(response.StatusCode, response)
}

func CErrResponse(c *gin.Context, statusCode int, customStatusCode int, errMessage string) {
	response := types.ApiResponse{
		StatusCode: customStatusCode,
		Error:      errMessage,
	}
	c.JSON(statusCode, response)
}

func SuccessResponse(c *gin.Context, message string, payload interface{}) {
	response := types.ApiResponse{
		StatusCode: http.StatusOK,
		Payload:    payload,
		Message:    message,
	}
	c.JSON(response.StatusCode, response)
}

func InternalServerError(c *gin.Context) {
	response := types.ApiResponse{
		StatusCode: http.StatusInternalServerError,
		Error:      "unexpected error occurred",
	}
	c.JSON(response.StatusCode, response)
}

func NewInternalServerError(c *gin.Context, format string, args ...interface{}) {
	logwrapper.Errorf(format, args...)
	response := types.ApiResponse{
		StatusCode: http.StatusInternalServerError,
		Error:      "unexpected error occurred",
	}
	c.JSON(response.StatusCode, response)
}
