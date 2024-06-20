package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// inicializa o router utilizando as configurações default do gin
	router := gin.Default()
	//inicia as rotas
	InitializeRoutes(router)
	//estamos rodando a nossa api
	// Get the port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the server
	router.Run("0.0.0.0:" + port)
}
