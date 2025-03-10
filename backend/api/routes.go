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

		api.POST("/ratings/employer", handlers.SaveEmployerRatings)

        api.GET("/head-data", handlers.GetHeadData)

        api.POST("/head-data", handlers.SaveHeadData)

		api.GET("/average", handlers.GetAverage)
	}
}