package models

type EmployerRatingDTO struct {
	UserEmail             string  `json:"userEmail"`
	RatingCardID          int     `json:"ratingCardId"`
	Question             string  `json:"question"`
	Category             string  `json:"category"`
	OrderID              int32   `json:"orderId"`
	TimeStampCandidate    *string `json:"timeStampCandidate"`
	RatingCandidate       *int    `json:"ratingCandidate"`
	TextResponseCandidate *string `json:"textResponseCandidate"`
	TimeStampEmployer    *string `json:"timeStampEmployer"`
	RatingEmployer        *int    `json:"ratingEmployer"`
	TextResponseEmployer *string `json:"textResponseEmployer"`
}
