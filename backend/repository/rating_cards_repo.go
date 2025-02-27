package repository

import (
	"backend/config"
	"backend/models"
	"context"
)

func GetRatingCards(ctx context.Context) ([]models.RatingCard, error) {
	var ratingCards []models.RatingCard

	result := config.DB.WithContext(ctx).Find(&ratingCards)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return ratingCards, nil
}