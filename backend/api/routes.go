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

		api.GET("/ratings/employer/candidates", handlers.GetRatingCandidates)

		api.GET("/ratings/employer", handlers.GetEmployerRatings)

		api.GET("/ratings/average", handlers.GetAverage)

		api.POST("/ratings/employer", handlers.SaveEmployerRatings)

        api.GET("/head-data", handlers.GetHeadData)

        api.POST("/head-data", handlers.SaveHeadData)

		api.GET("/management/agreed-candidates", handlers.GetAgreedCandidates)

		api.GET("/management/summary", handlers.GetManagementSummary)

		api.GET("/management/average", handlers.GetManagementAverage)
	}
}