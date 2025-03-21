package services

import (
	"backend/models"
	"backend/repository"
	"context"
	"sort"
	"time"
	"log"
)

func GetEmployerRatings(ctx context.Context, userEmail string) ([]models.EmployerRatingDTO, error) {
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
			dto.NotApplicableCandidate = rating.NotApplicableCandidate
	
			dto.TimeStampEmployer = nil
			if rating.TimeStampEmployer != nil && !rating.TimeStampEmployer.IsZero() {
				timeStrEmploy := rating.TimeStampEmployer.Format("2006-01-02 15:04:05")
				dto.TimeStampEmployer = &timeStrEmploy
			}
			dto.RatingEmployer = &rating.RatingEmployer
			dto.TextResponseEmployer = &rating.TextResponseEmployer
			dto.NotApplicableEmployer = rating.NotApplicableEmployer
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

func SaveEmployerRatings(ctx context.Context, candidateRatingsEmployer []models.EmployerRatingDTO) error {
	for _, dto := range candidateRatingsEmployer {
		currentTime := time.Now()

		rating := models.Rating {
			UserEmail: dto.UserEmail,
			RatingCardID: dto.RatingCardID,
			RatingEmployer: *dto.RatingEmployer,
			TextResponseEmployer: *dto.TextResponseEmployer,
			NotApplicableEmployer: dto.NotApplicableEmployer,
			TimeStampEmployer: &currentTime,
		}

		err := repository.SaveOrUpdateRating(ctx, &rating, false)
		if err != nil {
			log.Println("❌ Error saving employer rating:", err)
			return err
		}
	}
	return nil
}

func GetSeniorCandidates(ctx context.Context) ([]string, error) {
	emails, err := repository.GetSeniorCandidates(ctx)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}
	
	return emails, nil
}

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
