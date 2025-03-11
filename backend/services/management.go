package services

import (
	"backend/models"
	"backend/repository"
	"context"
	"fmt"
)

func GetHeadDataWithAgreement(ctx context.Context) ([]models.HeadDataDTO, error) {
	headDataSets, err := repository.GetHeadDataWithAgreement(ctx)
	if err != nil {
		return []models.HeadDataDTO{}, err
	}


	var headDataWithAgreement []models.HeadDataDTO
	for _, headData := range headDataSets {

		dto := models.HeadDataDTO{
			UserEmail:       headData.UserEmail,
			Name:            headData.Name,
			ExperienceSince: headData.ExperienceSince,
			StartAtProdyna:  headData.StartAtProdyna,
			Age:             headData.Age,
			Abstract:        headData.Abstract,
			AgreedOn:        headData.AgreedOn,
		}
		headDataWithAgreement = append(headDataWithAgreement, dto)
	}

	return headDataWithAgreement, nil
}

func GetManagementSummary(ctx context.Context, userEmail string) (models.ManagementSummaryDTO, error) {
	// Fetch the ratings for the specific user
	ratings, err := repository.GetRatingsByUserEmail(ctx, userEmail)
	if err != nil {
		return models.ManagementSummaryDTO{}, err
	}

	ratingCards, err := repository.GetRatingCards(ctx)
	if err != nil {
		return models.ManagementSummaryDTO{}, err
	}

	// Map to store aggregated rating data per category
	ratingMap := make(map[models.CategoryEnum][]int)

	// Group ratings by category
	for _, rating := range ratings {
		// Use GetCategoryByRatingCardID to fetch the category as CategoryEnum
		category, err := GetCategoryByRatingCardID(ratingCards, rating.RatingCardID)
		if err != nil {
			return models.ManagementSummaryDTO{}, err
		}

		// Convert category (string) to CategoryEnum
		categoryEnum := models.CategoryEnum(category)

		// Append the rating to the respective category
		ratingMap[categoryEnum] = append(ratingMap[categoryEnum], rating.RatingEmployer)
	}

	// Prepare the ManagementSummaryRatingDTO array
	var managementRatingsSummary []models.ManagementSummaryRatingDTO
	for category, ratings := range ratingMap {
		if len(ratings) == 0 { // Safety check to prevent division by zero
			continue
		}

		total := 0
		for _, r := range ratings {
			total += r
		}
		avg := float64(total) / float64(len(ratings)) // Calculate the average rating for the category

		// Add the category and average rating to the summary
		managementRatingsSummary = append(managementRatingsSummary, models.ManagementSummaryRatingDTO{
			Category: category, // CategoryEnum
			Rating:   avg,      // Average rating for the category
		})
	}

	// Return the ManagementSummaryDTO with the userEmail and ratings summary
	dto := models.ManagementSummaryDTO{
		UserEmail:                userEmail,
		ManagementRatingSummary:  managementRatingsSummary,
	}

	return dto, nil
}


func GetManagementAverage(ctx context.Context) ([]models.ManagementAverageDTO, error) {
	userEmailsWithAgreement, err := repository.GetUserEmailsWithAgreement(ctx)
	if err != nil {
		return nil, err
	}

	// Fetch all ratings for the given user emails
	ratings, err := repository.GetRatingsByEmails(ctx, userEmailsWithAgreement)
	if err != nil {
		return nil, err
	}

	ratingCards, err := repository.GetRatingCards(ctx)
	if err != nil {
		return nil, err
	}

	// Map to store aggregated rating data per category
	ratingMap := make(map[models.CategoryEnum][]int) // Change map key type to CategoryEnum

	// Group ratings by ratingCardId
	for _, rating := range ratings {
		// Use GetCategoryByRatingCardID to fetch the category as CategoryEnum
		category, err := GetCategoryByRatingCardID(ratingCards, rating.RatingCardID)
		if err != nil {
			return nil, err
		}

		// Convert category (string) to CategoryEnum
		categoryEnum := models.CategoryEnum(category) // Convert string to CategoryEnum

		ratingMap[categoryEnum] = append(ratingMap[categoryEnum], rating.RatingEmployer)
	}

	// Prepare the result list
	var averageRatings []models.ManagementAverageDTO
	for category, ratings := range ratingMap {
		if len(ratings) == 0 { // Safety check to prevent division by zero
			continue
		}

		total := 0
		for _, r := range ratings {
			total += r
		}
		avg := float64(total) / float64(len(ratings)) // Convert total to float64 for division

		averageRatings = append(averageRatings, models.ManagementAverageDTO{
			Category: category, // Now category is of type models.CategoryEnum
			Average:  avg,
		})
	}

	return averageRatings, nil
}

func GetCategoryByRatingCardID(ratingCards []models.RatingCard, ratingCardID int) (string, error) {
	for _, ratingCard := range ratingCards {
		if ratingCard.ID == ratingCardID {
			// Return the category of the found RatingCard
			return ratingCard.Category, nil
		}
	}

	// Return an error if no RatingCard with the specified ID is found
	return "", fmt.Errorf("RatingCard with ID %d not found", ratingCardID)
}