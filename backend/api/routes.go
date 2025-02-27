package api

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/rating-cards", handlers.GetRatingCards)
		
		api.POST("/ratings/save", handlers.SaveRatingRequest)
	}
}
