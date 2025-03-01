package services

import (
	"backend/models"
	"backend/repository"
	"context"
	"sort"
)

func GetCandidateRatingsForEmployer(ctx context.Context, userEmail string) ([]models.EmployerRatingDTO, error) {
    ratingCards, err := repository.GetRatingCards(ctx)
	if err != nil {
		return nil, err
	}

	var employerRatings []models.EmployerRatingDTO

	for _, card := range ratingCards {
		rating, err := repository.GetUserRating(ctx, userEmail, card.ID)
	
		dto := models.EmployerRatingDTO{
			UserEmail:    userEmail,
			RatingCardID: card.ID,
			Question:     card.Question,
			Category:     card.Category,
			OrderID:      card.OrderID,
		}
	
		if err == nil && rating != nil {
			dto.TimeStampCandidate = nil
			if rating.TimeStampCandidate != nil && !rating.TimeStampCandidate.IsZero() {
				timeStr := rating.TimeStampCandidate.Format("2006-01-02 15:04:05")
				dto.TimeStampCandidate = &timeStr
			}
			dto.RatingCandidate = &rating.RatingCandidate
			dto.TextResponseCandidate = &rating.TextResponseCandidate
	
			dto.TimeStampEmployer = nil
			if rating.TimeStampEmployer != nil && !rating.TimeStampEmployer.IsZero() {
				timeStrEmploy := rating.TimeStampEmployer.Format("2006-01-02 15:04:05")
				dto.TimeStampEmployer = &timeStrEmploy
			}
			dto.RatingEmployer = &rating.RatingEmployer
			dto.TextResponseEmployer = &rating.TextResponseEmployer
		}
	
		employerRatings = append(employerRatings, dto)
	}	

	sort.Slice(employerRatings, func(i, j int) bool {
		if employerRatings[i].Category == employerRatings[j].Category {
			return employerRatings[i].OrderID < employerRatings[j].OrderID
		}
		return employerRatings[i].Category < employerRatings[j].Category
	})

	return employerRatings, nil
}
