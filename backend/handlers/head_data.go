package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHeadData godoc
// @Summary Get head data for candidate
// @Description Fetches head data for certain candidate
// @Tags head-data
// @Produce json
// @Param userEmail query string true "User Email"
// @Success 200 {object} models.HeadDataDTO
// @Failure 500 {object} models.ErrorResponse
// @Router /head-data [get]
func GetHeadData(c *gin.Context) {

	userEmail := c.Query("userEmail")
	if userEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userEmail is required"})
		return
	}

	headData, err := services.GetHeadData(c.Request.Context(), userEmail)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch headData"})
		return
	}

	c.JSON(http.StatusOK, headData)
}

// SaveHeadData godoc
// @Summary Save head data
// @Description Stores or updates head data in the database
// @Tags head-data
// @Accept json
// @Produce json
// @Param ratings body []models.HeadDataDTO true "head data"
// @Success 201 {object} map[string]string "message: HeadData saved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request payload"
// @Failure 500 {object} models.ErrorResponse "Failed to save headData"
// @Router /head-data [post]
func SaveHeadData(c *gin.Context) {
	var headDataDTO models.HeadDataDTO
	if err := c.ShouldBindJSON(&headDataDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := services.SaveHeadData(c.Request.Context(), headDataDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save ratings"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ratings saved successfully"})
}
