package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize router
	router := gin.Default()

	// Routes
	initializeRouter(router)

	// Running the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
