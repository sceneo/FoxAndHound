package repository

import (
	"backend/config"
	"backend/models"
	"fmt"
	"time"
)

func SaveCandidateRatings(ratingRequest models.RatingRequest) error {
	timestamp, err := time.Parse(time.RFC3339, ratingRequest.TimeStamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp format: %v", err)
	}

	tx := config.DB.Begin()

	for _, rating := range ratingRequest.Ratings {
		rating.TimeStampCandidate = &timestamp
		rating.UserEmail = ratingRequest.UserEmail

		if err := tx.Create(&rating).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("error saving rating: %v", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}