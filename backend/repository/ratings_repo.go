package repository

import (
	"backend/config"
	"backend/models"
	"context"
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

func SaveCandidateRatings(ctx context.Context, candidateRatings []models.CandidateRatingDTO) error {
	tx := config.DB.WithContext(ctx).Begin()

	for _, dto := range candidateRatings {
		var rating models.Rating
		result := tx.Where("user_email = ? AND rating_card_id = ?", dto.UserEmail, dto.RatingCardID).First(&rating)

		currentTime := time.Now()

		if result.RowsAffected > 0 {
			log.Printf("🔄 Updating rating for user: %s, card: %d\n", dto.UserEmail, dto.RatingCardID)

			err := tx.Model(&rating).Updates(models.Rating{
				RatingCandidate:     *dto.RatingCandidate,
				TextResponseCandidate: *dto.TextResponseCandidate,
				TimeStampCandidate:  &currentTime,
			}).Error

			if err != nil {
				tx.Rollback()
				log.Println("❌ Error updating rating:", err)
				return err
			}

		} else {
			log.Printf("🆕 Creating new rating for user: %s, card: %d\n", dto.UserEmail, dto.RatingCardID)

			newRating := models.Rating{
				UserEmail:            dto.UserEmail,
				RatingCardID:         dto.RatingCardID,
				RatingCandidate:      *dto.RatingCandidate,
				TextResponseCandidate: *dto.TextResponseCandidate,
				TimeStampCandidate:   &currentTime,
			}

			err := tx.Create(&newRating).Error
			if err != nil {
				tx.Rollback()
				log.Println("❌ Error creating new rating:", err)
				return err
			}
		}
	}

	return tx.Commit().Error
}