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

// GetCandidateRatings godoc
// @Summary Get candidate ratings
// @Description Fetches rating cards and enriches them with existing ratings for a given user
// @Tags rating
// @Produce json
// @Param userEmail query string true "User Email"
// @Success 200 {array} models.CandidateRatingDTO
// @Failure 500 {object} models.ErrorResponse
// @Router /ratings/candidate [get]
func GetCandidateRatings(c *gin.Context) {
	userEmail := c.Query("userEmail")
	if userEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userEmail is required"})
		return
	}

	ratings, err := repository.GetCandidateRatings(c.Request.Context(), userEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch candidate ratings"})
		return
	}

	c.JSON(http.StatusOK, ratings)
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
