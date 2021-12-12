package httphelper

import (
	"net/http"
	"netsepio-api/types"

	"github.com/gin-gonic/gin"
)

func ErrResponse(c *gin.Context, statusCode int, errMessage string) {
	response := types.ApiResponse{
		Status: statusCode,
		Error:  errMessage,
	}
	c.JSON(response.Status, response)
}

func SuccessResponse(c *gin.Context, message string, payload interface{}) {
	response := types.ApiResponse{
		Status:  http.StatusOK,
		Payload: payload,
		Message: message,
	}
	c.JSON(response.Status, response)
}
