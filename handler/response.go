package handler

import (
	"fmt"
	"net/http"

	"github.com/BrunoBerval/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, msg string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"message": msg,
		"code":    code,
	})
}

func sendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-type", "aplication/jason")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}
