package models

import (
	"time"
)

type RatingCard struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Question string `json:"question"`
	Category string `json:"category"`
	OrderID  int32  `json:"orderId"`
}

type Rating struct {
	ID                    int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserEmail             string       `json:"userEmail"`
	TimeStampCandidate    *time.Time `json:"timeStampCandidate" gorm:"default:null column:timestamp_candidate"`
	TimeStampEmployer     *time.Time `json:"timeStampEmployer" gorm:"default:null column:timestamp_employer"`
	RatingCardID          int       `json:"ratingCardId" gorm:"foreignKey:RatingCardID"`
	RatingCandidate       int       `json:"ratingCandidate"`
	TextResponseCandidate string    `json:"textResponseCandidate"`
	RatingEmployer        int       `json:"ratingEmployer"`
	TextResponseEmployer  string    `json:"textResponseEmployer"`
}

// RatingRequest struct for handling JSON input
type RatingRequest struct {
	UserEmail    string      `json:"userEmail"`
	TimeStamp string   `json:"timeStamp"`
	Ratings   []Rating `json:"ratings"`
}