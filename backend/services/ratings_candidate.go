package services

import (
	"backend/models"
	"backend/repository"
	"context"
	"log"
	"sort"
	"time"
)

func GetCandidateRatings(ctx context.Context, userEmail string) ([]models.CandidateRatingDTO, error) {
	ratingCards, err := repository.GetRatingCards(ctx)
	if err != nil {
		return nil, err
	}

	var candidateRatings []models.CandidateRatingDTO

	for _, card := range ratingCards {
		rating, err := repository.GetUserRating(ctx, userEmail, card.ID)

		dto := models.CandidateRatingDTO{
			UserEmail:   userEmail,
			RatingCardID: card.ID,
			Question:    card.Question,
			Category:    card.Category,
			OrderID:     card.OrderID,
		}

		if err == nil && rating != nil {
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
	for _, dto := range candidateRatings {
		currentTime := time.Now()

		rating := models.Rating{
			UserEmail:            dto.UserEmail,
			RatingCardID:         dto.RatingCardID,
			RatingCandidate:      *dto.RatingCandidate,
			TextResponseCandidate: *dto.TextResponseCandidate,
			TimeStampCandidate:   &currentTime,
		}

		err := repository.SaveOrUpdateCandidateRating(ctx, &rating)
		if err != nil {
			log.Println("❌ Error saving candidate rating:", err)
			return err
		}
	}
	return nil
}
