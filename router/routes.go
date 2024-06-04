package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Define uma rota
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(c *gin.Context) {
			// Adicione aqui a lógica do handler
			c.JSON(http.StatusOK, gin.H{
				"message": "GET opening",
			})
		})

		v1.POST("/opening", func(c *gin.Context) {
			// Adicione aqui a lógica do handler
			c.JSON(http.StatusOK, gin.H{
				"message": "POST opening",
			})
		})

		v1.DELETE("/opening", func(c *gin.Context) {
			// Adicione aqui a lógica do handler
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE opening",
			})
		})

		v1.PUT("/opening", func(c *gin.Context) {
			// Adicione aqui a lógica do handler
			c.JSON(http.StatusOK, gin.H{
				"message": "PUT opening",
			})
		})

		v1.GET("/openings", func(c *gin.Context) {
			// Adicione aqui a lógica do handler
			c.JSON(http.StatusOK, gin.H{
				"message": "GET openings",
			})
		})
	}
}
