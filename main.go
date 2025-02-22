package main

import (
	"github.com/MH0R/Gopportunities/config"
	_ "github.com/MH0R/Gopportunities/docs"
	"github.com/MH0R/Gopportunities/router"
)

var (
	logger config.Logger
)

//	@title			Gopportunities
//	@version		1.0
//	@description	CRUD using Gin, Swagger, GoORM and SQLite
//	@contact.name	github.com/MH0R
//	@host			localhost:8080
//	@BasePath		/api/v1

func main() {
	logger = *config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errf("Error initializing config: %v", err)
		return
	}
	router.Initialize()
}
