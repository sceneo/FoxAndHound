package handlers

import (
	"backend/repository"
	"net/http"
	"backend/models"

	"github.com/gin-gonic/gin"
)

// GetRatingCards godoc
// @Summary Get all rating cards
// @Description Retrieves all rating cards from the database
// @Tags rating
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

// POST /api/ratings/save
func SaveRatingRequest(c *gin.Context) {
	var ratingRequest models.RatingRequest

	// Parse JSON body
	if err := c.ShouldBindJSON(&ratingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Save candidate ratings
	err := repository.SaveCandidateRatings(ratingRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save candidate ratings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ratings saved successfully"})
}
