package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRatingCandidates godoc
// @Summary Get rating candidates
// @Description Fetches rating candidates
// @Tags rating-employer
// @Produce json
// @Success 200 {array} string
// @Failure 500 {object} models.ErrorResponse
// @Router /ratings/employer/candidates [get]
func GetRatingCandidates(c *gin.Context) {
	candidates, err := services.GetSeniorCandidates(c.Request.Context())
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch candidates"})
		return
	}

	c.JSON(http.StatusOK, candidates)
}

// GetEmployerRatings godoc
// @Summary Get candidate ratings for employer
// @Description Fetches rating cards and enriches them with existing ratings for a given user
// @Tags rating-employer
// @Produce json
// @Param userEmail query string true "User Email"
// @Success 200 {array} models.EmployerRatingDTO
// @Failure 500 {object} models.ErrorResponse
// @Router /ratings/employer [get]
func GetEmployerRatings(c *gin.Context) {
	userEmail := c.Query("userEmail")
	if userEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userEmail is required"})
		return
	}

	ratings, err := services.GetEmployerRatings(c.Request.Context(), userEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch candidate ratings"})
		return
	}

	c.JSON(http.StatusOK, ratings)
}

// SaveEmployerRatings godoc
// @Summary Save employer ratings
// @Description Stores or updates candidate ratings of employer in the database
// @Tags rating-employer
// @Accept json
// @Produce json
// @Param ratings body []models.EmployerRatingDTO true "List of candidate ratings"
// @Success 201 {object} map[string]string "message: Ratings saved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request payload"
// @Failure 500 {object} models.ErrorResponse "Failed to save ratings"
// @Router /ratings/employer [post]
func SaveEmployerRatings(c *gin.Context) {
	var ratings []models.EmployerRatingDTO
	if err := c.ShouldBindJSON(&ratings); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

	err := services.SaveEmployerRatings(c.Request.Context(), ratings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save ratings"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ratings saved successfully"})
}