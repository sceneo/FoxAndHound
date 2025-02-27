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

// SaveCandidateRatings godoc
// @Summary Save candidate ratings
// @Description Stores or updates candidate ratings in the database
// @Tags rating
// @Accept json
// @Produce json
// @Param ratings body []models.CandidateRatingDTO true "List of candidate ratings"
// @Success 201 {object} map[string]string "message: Ratings saved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request payload"
// @Failure 500 {object} models.ErrorResponse "Failed to save ratings"
// @Router /ratings/candidate [post]
func SaveCandidateRatings(c *gin.Context) {
	var ratings []models.CandidateRatingDTO
    if err := c.ShouldBindJSON(&ratings); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    err := repository.SaveCandidateRatings(c.Request.Context(), ratings)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save ratings"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Ratings saved successfully"})
}
