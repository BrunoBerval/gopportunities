package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(c *gin.Context) {
	// Adicione aqui a lógica do handler
	c.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func CreateOpeningHandler(c *gin.Context) {
	// Adicione aqui a lógica do handler
	c.JSON(http.StatusOK, gin.H{
		"message": "POST opening",
	})
}

func DeleteOpeningHandler(c *gin.Context) {
	// Adicione aqui a lógica do handler
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}

func UpdateOpeningHandler(c *gin.Context) {
	// Adicione aqui a lógica do handler
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}

func ListOpeningsHandler(c *gin.Context) {
	// Adicione aqui a lógica do handler
	c.JSON(http.StatusOK, gin.H{
		"message": "GET openings",
	})
}
