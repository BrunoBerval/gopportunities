package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// inicializa o router utilizando as configurações default do gin
	router := gin.Default()
	//define uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//estamos rodando a nossa api
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
