package repository

import (
	"backend/config"
	"backend/models"
	"log"
)

func GetRatingCards() ([]models.RatingCard, error) {
	var ratingCards []models.RatingCard
	result := config.DB.Find(&ratingCards)
	if result.Error != nil {
		log.Println("❌ Error fetching rating cards:", result.Error)
		return nil, result.Error
	}
	return ratingCards, nil
}
