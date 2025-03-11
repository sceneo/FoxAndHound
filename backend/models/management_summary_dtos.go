package models

type ManagementSummaryDTO struct {
	UserEmail               string                       `json:"userEmail"`
	ManagementRatingSummary []ManagementSummaryRatingDTO `json:"ratings"` // Array of management ratings
}
