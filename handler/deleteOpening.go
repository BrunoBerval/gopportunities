package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(c *gin.Context) {
	// Adicione aqui a lógica do handler
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}
