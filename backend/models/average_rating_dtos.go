package models

type AverageRatingDTO struct {
	RatingCardID          int     `json:"ratingCardId"`
	NumberOfAgreedRatings int     `json:"NumberOfAgreedRatings"`
	Average               float64 `json:"average"`
}
