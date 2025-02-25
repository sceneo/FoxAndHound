package models

import (
	"time"
	"gorm.io/gorm"
)

type RatingCard struct {
	gorm.Model
	ID       string `json:"id" gorm:"primaryKey"`
	Question string `json:"question"`
	Category string `json:"category"`
	OrderId  int32  `json:"orderId"`
}

type Rating struct {
	gorm.Model
	UserId                int       `json:"userId"`
	TimeStamp             time.Time `json:"timeStamp"`
	RatingCardId          int       `json:"ratingCardId"`
	RatingCandidate       int       `json:"ratingCandidate"`
	TextResponseCandidate string    `json:"textResponseCandidate"`
}

// RatingRequest struct for handling JSON input
type RatingRequest struct {
	UserId    int      `json:"userId"`
	TimeStamp string   `json:"timeStamp"`
	Ratings   []Rating `json:"ratings"`
}

// TotalRating model for aggregated results
type TotalRating struct {
	gorm.Model
	RatingCardId     int     `json:"ratingCardId"`
	AverageRating    float64 `json:"averageRating"`
	TotalSubmissions int     `json:"totalSubmissions"`
}