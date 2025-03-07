package services

import (
	"backend/models"
	"backend/repository"
	"context"
	"log"
)

func GetHeadData(ctx context.Context, userEmail string) (models.HeadDataDTO, error) {
	headData, err := repository.GetHeadData(ctx, userEmail)
	if err != nil {
		return models.HeadDataDTO{}, err
	}

	dto := models.HeadDataDTO{
		UserEmail:       userEmail,
		Name:            headData.Name,
		ExperienceSince: headData.ExperienceSince,
		StartAtProdyna:  headData.StartAtProdyna,
		Age:             headData.Age,
		Abstract:        headData.Abstract,
		AgreedOn:        headData.AgreedOn,
	}

	return dto, nil
}

func SaveHeadData(ctx context.Context, headDataDTO models.HeadDataDTO) error {
	headData := models.HeadData{
		UserEmail:       headDataDTO.UserEmail,
		Name:            headDataDTO.Name,
		ExperienceSince: headDataDTO.ExperienceSince,
		StartAtProdyna:  headDataDTO.StartAtProdyna,
		Age:             headDataDTO.Age,
		Abstract:        headDataDTO.Abstract,
		AgreedOn:        headDataDTO.AgreedOn,
	}

	err := repository.SaveOrUpdateHeadData(ctx, headData)
	if err != nil {
		log.Println("❌ Error saving head data:", err)
		return err
	}
	return nil
}
