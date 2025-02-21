package router

import (
	"github.com/MH0R/Gopportunities/router/handler"
	"github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine) {
	// Routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.CreateOpeningHandler)

		v1.POST("/opening", handler.ShowOpeningHandler)

		v1.PUT("/opening", handler.DeleteOpeningHandler)

		v1.DELETE("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)

	}
}
