package handlers

import (
	"backend/repository"
	"backend/services"
	"backend/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetRatingCards godoc
// @Summary Get all rating cards
// @Description Retrieves all rating cards from the database
// @Tags rating-card
// @Produce json
// @Success 200 {array} models.RatingCard
// @Failure 500 {object} models.ErrorResponse
// @Router /rating-cards [get]
func GetRatingCards(c *gin.Context) {
	data, err := repository.GetRatingCards(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch rating cards"})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetAverage godoc
// @Summary Get average per card
// @Description Fetches average for every card - agreed on flag is a must
// @Tags head-data
// @Produce json
// @Success 200 {array} models.AverageRatingDTO
// @Failure 500 {object} models.ErrorResponse
// @Router /ratings/average [get]
func GetAverage(c *gin.Context) {

	averageRatings, err := services.GetAverageRatings(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch headData"})
		return
	}

	c.JSON(http.StatusOK, averageRatings)
}
