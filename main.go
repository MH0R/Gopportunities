package main

import (
	"github.com/MH0R/Gopportunities/config"
	"github.com/MH0R/Gopportunities/router"
)

var (
	logger config.Logger
)

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
