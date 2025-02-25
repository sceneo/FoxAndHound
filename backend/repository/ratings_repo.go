package repository

import (
	"backend/config"
	"backend/models"
	"fmt"
	"log"
	"time"
)

func SaveCandidateRatings(ratingRequest models.RatingRequest) error {
	timestamp, err := time.Parse(time.RFC3339, ratingRequest.TimeStamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp format: %v", err)
	}

	tx := config.DB.Begin()

	for _, rating := range ratingRequest.Ratings {
		rating.TimeStamp = timestamp
		rating.UserId = ratingRequest.UserId

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

func SaveTotalRating() error {
	var totalRatings []models.TotalRating

	rows, err := config.DB.Raw(`
		SELECT RatingCardId, AVG(RatingCandidate) AS AverageRating, COUNT(*) AS TotalSubmissions
		FROM ratings
		GROUP BY RatingCardId
	`).Rows()
	if err != nil {
		return fmt.Errorf("error fetching ratings: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var total models.TotalRating
		if err := rows.Scan(&total.RatingCardId, &total.AverageRating, &total.TotalSubmissions); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		totalRatings = append(totalRatings, total)
	}

	for _, total := range totalRatings {
		err := config.DB.
			Where("rating_card_id = ?", total.RatingCardId).
			Assign(models.TotalRating{
				AverageRating:    total.AverageRating,
				TotalSubmissions: total.TotalSubmissions,
			}).
			FirstOrCreate(&total).Error

		if err != nil {
			log.Println("❌ Error saving total rating:", err)
			return err
		}
	}

	return nil
}
