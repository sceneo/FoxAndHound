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
	ID                     int        `json:"id" gorm:"primaryKey;autoIncrement"`
	UserEmail              string     `json:"userEmail"`
	TimeStampCandidate     *time.Time `json:"timeStampCandidate" gorm:"default:null"`
	TimeStampEmployer      *time.Time `json:"timeStampEmployer" gorm:"default:null"`
	RatingCardID           int        `json:"ratingCardId" gorm:"foreignKey:RatingCardID"`
	RatingCandidate        int        `json:"ratingCandidate"`
	TextResponseCandidate  string     `json:"textResponseCandidate"`
	NotApplicableCandidate bool       `json:"notApplicableCandidate"`
	RatingEmployer         int        `json:"ratingEmployer"`
	TextResponseEmployer   string     `json:"textResponseEmployer"`
	NotApplicableEmployer  bool       `json:"notApplicableEmployer"`
	IsClosed               bool       `json:"isClosed"`
}
