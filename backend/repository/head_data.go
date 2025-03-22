package repository

import (
	"backend/config"
	"backend/models"
	"context"
	"log"
)

func GetHeadData(ctx context.Context, userEmail string) (*models.HeadData, error) {
	var headData models.HeadData
	result := config.DB.WithContext(ctx).
		Where("user_email = ?", userEmail).
		First(&headData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &headData, nil
}

func GetHeadDataWithAgreement(ctx context.Context) ([]models.HeadData, error) {
	var headData []models.HeadData
	result := config.DB.WithContext(ctx).
		Model(&models.HeadData{}).
		Where("agreed_on = ?", true).
		Find(&headData)

	if result.Error != nil {
		return nil, result.Error
	}

	return headData, nil
}

func GetUserEmailsWithAgreement(ctx context.Context) ([]string, error) {
	var userEmails []string
	result := config.DB.WithContext(ctx).
		Model(&models.HeadData{}).
		Where("agreed_on = ?", true).
		Pluck("user_email", &userEmails) // Extract only user_email column

	if result.Error != nil {
		return nil, result.Error
	}

	return userEmails, nil
}

func SaveOrUpdateHeadData(ctx context.Context, headData models.HeadData) error {
	tx := config.DB.WithContext(ctx).Begin()

	var existing models.HeadData
	result := tx.Where("user_email = ?", headData.UserEmail).First(&existing)

	if result.RowsAffected > 0 {
		updateFields := map[string]interface{}{
			"name":                 headData.Name,
			"experience_since":     headData.ExperienceSince,
			"start_at_prodyna":     headData.StartAtProdyna,
			"age":                  headData.Age,
			"abstract":             headData.Abstract,
			"agreed_on":            headData.AgreedOn,
			"submit_to_management": headData.SubmitToManagement,
			"country":              headData.Country,
			"is_promoted":          headData.IsPromoted,
		}

		err := tx.Model(&existing).Where("user_email = ?", headData.UserEmail).Updates(updateFields).Error
		if err != nil {
			tx.Rollback()
			log.Println("❌ Error updating head-data:", err)
			return err
		}
	} else {
		err := tx.Create(&headData).Error
		if err != nil {
			tx.Rollback()
			log.Println("❌ Error creating head-data:", err)
			return err
		}
	}

	return tx.Commit().Error
}
