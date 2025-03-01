package repository

import (
	"backend/config"
	"backend/models"
	"context"
	"log"
)

func GetUserRating(ctx context.Context, userEmail string, ratingCardID int) (*models.Rating, error) {
	var rating models.Rating
	result := config.DB.WithContext(ctx).
		Where("user_email = ? AND rating_card_id = ?", userEmail, ratingCardID).
		First(&rating)

	if result.Error != nil {
		return nil, result.Error
	}

	return &rating, nil
}

func SaveOrUpdateCandidateRating(ctx context.Context, rating *models.Rating) error {
	tx := config.DB.WithContext(ctx).Begin()

	var existing models.Rating
	result := tx.Where("user_email = ? AND rating_card_id = ?", rating.UserEmail, rating.RatingCardID).First(&existing)

	if result.RowsAffected > 0 {
		err := tx.Model(&existing).Updates(rating).Error
		if err != nil {
			tx.Rollback()
			log.Println("❌ Error updating rating:", err)
			return err
		}
	} else {
		err := tx.Create(rating).Error
		if err != nil {
			tx.Rollback()
			log.Println("❌ Error creating rating:", err)
			return err
		}
	}

	return tx.Commit().Error
}