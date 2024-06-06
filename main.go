package main

import (
	"github.com/BrunoBerval/gopportunities/config"
	"github.com/BrunoBerval/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger()
	// Iniciar as configurações
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	//Iniciar o router
	router.Initialize()
}
