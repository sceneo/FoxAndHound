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

func SaveOrUpdateHeadData(ctx context.Context, headData models.HeadData) error {
	tx := config.DB.WithContext(ctx).Begin()

	var existing models.HeadData
	result := tx.Where("user_email = ?", headData.UserEmail).First(&existing)


    if result.RowsAffected > 0 {
            updateFields := map[string]interface{}{}

            updateFields["user_email"] = headData.UserEmail
            updateFields["name"] = headData.Name
            updateFields["experience_since"] = headData.ExperienceSince
            updateFields["start_at_prodyna"] = headData.StartAtProdyna
            updateFields["age"] = headData.Age
            updateFields["abstract"] = headData.Abstract

            err := tx.Model(&existing).Updates(updateFields).Error

        if err != nil {
            tx.Rollback()
            log.Println("❌ Error updating rating:", err)
            return err
        }
    } else {
        err := tx.Create(headData).Error

        if err != nil {
            tx.Rollback()
            log.Println("❌ Error creating rating:", err)
            return err
        }
    }

	return tx.Commit().Error
}