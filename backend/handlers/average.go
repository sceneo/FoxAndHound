package handlers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAverage godoc
// @Summary Get average per card
// @Description Fetches average for every card - agreed on flag is a must
// @Tags head-data
// @Produce json
// @Success 200 {array} string
// @Failure 500 {object} models.ErrorResponse
// @Router /average [get]
func GetAverage(c *gin.Context) {

	averageRatings, err := services.GetAverageRatings(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch headData"})
		return
	}

	c.JSON(http.StatusOK, averageRatings)
}

