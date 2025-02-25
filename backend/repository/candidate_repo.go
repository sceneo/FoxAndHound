package repository

import (
	"backend/config"
	"backend/models"
	"log"
)

func GetAllCandidates() ([]models.Candidate, error) {
	var candidates []models.Candidate
	result := config.DB.Table("Rating").Distinct("UserId").Find(&candidates)
	if result.Error != nil {
		log.Println("❌ Error fetching candidates:", result.Error)
		return nil, result.Error
	}
	return candidates, nil
}
