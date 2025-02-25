package handlers

import (
	"backend/models"
	"backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/rating-cards
func GetRatingCards(c *gin.Context) {
	data, err := repository.GetRatingCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rating cards"})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GET /api/candidates
func GetAllCandidates(c *gin.Context) {
	data, err := repository.GetAllCandidates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch candidates"})
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

	// Save total ratings
	err = repository.SaveTotalRating()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save total ratings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ratings saved successfully"})
}
