package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	r := gin.Default()

	// Routes
	initializeRouter(r)

	// Running the server
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
