package services

import (
	"backend/models"
	"backend/repository"
	"context"
)

func GetAverageRatings(ctx context.Context) ([]models.AverageRatingDTO, error) {
	userEmailsWithAgreement, err := repository.GetUserEmailsWithAgreement(ctx)
	if err != nil {
		return nil, err
	}

	// Fetch all ratings for the given user emails
	ratings, err := repository.GetRatingsByEmails(ctx, userEmailsWithAgreement)
	if err != nil {
		return nil, err
	}

	// Map to store aggregated rating data per ratingCardId
	ratingMap := make(map[int][]int)

	// Group ratings by ratingCardId
	for _, rating := range ratings {
		ratingMap[rating.RatingCardID] = append(ratingMap[rating.RatingCardID], rating.RatingEmployer)
	}

	// Prepare the result list
	var averageRatings []models.AverageRatingDTO
	for ratingCardID, ratings := range ratingMap {
		if len(ratings) == 0 { // Safety check to prevent division by zero
			continue
		}

		total := 0
		for _, r := range ratings {
			total += r
		}
		avg := float64(total) / float64(len(ratings)) // Convert total to float64 for division

		averageRatings = append(averageRatings, models.AverageRatingDTO{
			RatingCardID:          ratingCardID,
			Average:               avg,
			NumberOfAgreedRatings: len(ratings),
		})
	}

	return averageRatings, nil
}
