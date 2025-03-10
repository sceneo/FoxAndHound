package repository

import (
	"backend/config"
	"backend/models"
	"context"
	"log"
)

func GetUserRating(ctx context.Context, userEmail string, ratingCardID int) (*models.Rating, error) {
	var rating models.Rating
	result := config.DB.WithContext(ctx).
		Where("user_email = ? AND rating_card_id = ?", userEmail, ratingCardID).
		First(&rating)

	if result.Error != nil {
		return nil, result.Error
	}

	return &rating, nil
}

func SaveOrUpdateRating(ctx context.Context, rating *models.Rating, isCandidate bool) error {
	tx := config.DB.WithContext(ctx).Begin()

	var existing models.Rating
	result := tx.Where("user_email = ? AND rating_card_id = ?", rating.UserEmail, rating.RatingCardID).First(&existing)

	if result.RowsAffected > 0 {
		updateFields := map[string]interface{}{}

		if isCandidate {
			updateFields["time_stamp_candidate"] = rating.TimeStampCandidate
			updateFields["rating_candidate"] = rating.RatingCandidate
			updateFields["text_response_candidate"] = rating.TextResponseCandidate
			updateFields["not_applicable_candidate"] = rating.NotApplicableCandidate
		} else {
			updateFields["time_stamp_employer"] = rating.TimeStampEmployer
			updateFields["rating_employer"] = rating.RatingEmployer
			updateFields["text_response_employer"] = rating.TextResponseEmployer
			updateFields["not_applicable_employer"] = rating.NotApplicableEmployer
		}

		err := tx.Model(&existing).Updates(updateFields).Error
		if err != nil {
			tx.Rollback()
			log.Println("❌ Error updating rating:", err)
			return err
		}
	} else {
		err := tx.Create(rating).Error
		if err != nil {
			tx.Rollback()
			log.Println("❌ Error creating rating:", err)
			return err
		}
	}

	return tx.Commit().Error
}

func GetRatingsByEmails(ctx context.Context, userEmails []string) ([]models.Rating, error) {
	var ratings []models.Rating
	result := config.DB.WithContext(ctx).
		Where("user_email IN ?", userEmails).
		Find(&ratings)

	if result.Error != nil {
		return nil, result.Error
	}
	return ratings, nil
}

func GetSeniorCandidates(ctx context.Context) ([]string, error) {
	var userEmails []string
	err := config.DB.WithContext(ctx).
		Model(&models.Rating{}).
		Distinct("user_email").
		Pluck("user_email", &userEmails).Error

	if err != nil {
		log.Println("❌ Error fetching distinct user emails:", err)
		return nil, err
	}

	return userEmails, nil
}
