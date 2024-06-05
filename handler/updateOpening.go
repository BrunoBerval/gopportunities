package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(c *gin.Context) {
	// Adicione aqui a lógica do handler
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}
