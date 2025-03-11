package models

type ManagementSummaryRatingDTO struct {
	Category CategoryEnum `json:"category"`
	Rating   float64      `json:"rating"`
}
