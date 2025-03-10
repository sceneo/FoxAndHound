package services

import (
	"backend/models"
	"backend/repository"
	"context"
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

	dto := models.ManagementSummaryDTO{
		UserEmail:       userEmail,
		// TODO: fill
	}

	return dto, nil
}