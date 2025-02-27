package api

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/rating-cards", handlers.GetRatingCards)

		api.GET("/ratings/candidate", handlers.GetCandidateRatings)
		
		api.POST("/ratings/candidate", handlers.SaveCandidateRatings)
	}
}
