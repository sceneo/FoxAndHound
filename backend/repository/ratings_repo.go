package repository

import (
	"backend/config"
	"backend/models"
	"context"
	"fmt"
	"log"
	"sort"
	"time"
)

func GetCandidateRatings(ctx context.Context, userEmail string) ([]models.CandidateRatingDTO, error) {
	var ratingCards []models.RatingCard
	var candidateRatings []models.CandidateRatingDTO

	if err := config.DB.WithContext(ctx).Find(&ratingCards).Error; err != nil {
		log.Println("❌ Error fetching rating cards:", err)
		return nil, err
	}

	for _, card := range ratingCards {
		var rating models.Rating
		result := config.DB.WithContext(ctx).Where("user_email = ? AND rating_card_id = ?", userEmail, card.ID).First(&rating)

		dto := models.CandidateRatingDTO{
			UserEmail:   userEmail,
			RatingCardID: card.ID,
			Question:    card.Question,
			Category:    card.Category,
			OrderID:     card.OrderID,
		}

		if result.RowsAffected > 0 {
			dto.TimeStampCandidate = nil
			if !rating.TimeStampCandidate.IsZero() {
				timeStr := rating.TimeStampCandidate.Format("2006-01-02 15:04:05")
				dto.TimeStampCandidate = &timeStr
			}
			dto.RatingCandidate = &rating.RatingCandidate
			dto.TextResponseCandidate = &rating.TextResponseCandidate
		}

		candidateRatings = append(candidateRatings, dto)
	}

	sort.Slice(candidateRatings, func(i, j int) bool {
		if candidateRatings[i].Category == candidateRatings[j].Category {
			return candidateRatings[i].OrderID < candidateRatings[j].OrderID
		}
		return candidateRatings[i].Category < candidateRatings[j].Category
	})

	return candidateRatings, nil
}

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