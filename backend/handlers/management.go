package handlers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAgreedCandidates godoc
// @Summary Get agreed candidates
// @Description Fetches head data for agreed candidates
// @Tags head-data
// @Produce json
// @Success 200 {array} models.HeadDataDTO
// @Failure 500 {object} models.ErrorResponse
// @Router /management/agreed-candidates [get]
func GetAgreedCandidates(c *gin.Context) {
	headData, err := services.GetHeadDataWithAgreement(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch headData"})
		return
	}

	c.JSON(http.StatusOK, headData)
}

// GetManagementSummary godoc
// @Summary Get management summary for candidate
// @Description Fetches management summary for certain candidate
// @Tags management-summary
// @Produce json
// @Param userEmail query string true "User Email"
// @Success 200 {object} models.ManagementSummaryDTO
// @Failure 500 {object} models.ErrorResponse
// @Router /management/summary [get]
func GetManagementSummary(c *gin.Context) {

	userEmail := c.Query("userEmail")
	if userEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userEmail is required"})
		return
	}

	managementSummary, err := services.GetManagementSummary(c.Request.Context(), userEmail)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch management summary"})
		return
	}

	c.JSON(http.StatusOK, managementSummary)
}

// GetManagementAverage godoc
// @Summary Get management average for cards
// @Description Fetches management average for cards
// @Tags management-average
// @Produce json
// @Success 200 {array} models.ManagementAverageDTO
// @Failure 500 {object} models.ErrorResponse
// @Router /management/average [get]
func GetManagementAverage(c *gin.Context) {

	managementAverage, err := services.GetManagementAverage(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch management average"})
		return
	}

	c.JSON(http.StatusOK, managementAverage)
}