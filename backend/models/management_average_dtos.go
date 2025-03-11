package models

type ManagementAverageDTO struct {
	Category CategoryEnum `json:"category"`
	Average  float64      `json:"average"`
}
