package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// inicializa o router utilizando as configurações default do gin
	router := gin.Default()
	//inicia as rotas
	InitializeRoutes(router)
	//estamos rodando a nossa api
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
