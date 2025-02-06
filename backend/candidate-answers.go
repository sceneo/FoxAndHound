package main

import (
	"database/sql"
	"fmt"
	"time"
)

// Rating holds the structure of a single rating entry.
type Rating struct {
	RatingCardId          int    `json:"ratingCardId"`
	RatingCandidate       int    `json:"ratingCandidate"`
	TextResponseCandidate string `json:"textResponseCandidate"`
}

// RatingRequest holds the overall structure for the rating request.
type RatingRequest struct {
	UserId    int      `json:"userId"`
	TimeStamp string   `json:"timeStamp"`
	Ratings   []Rating `json:"ratings"`
}

// SaveCandidateRatings inserts the ratings into the database for candidates.
func SaveCandidateRatings(db *sql.DB, ratingRequest RatingRequest) error {
	// Parse timeStamp to time.Time format
	timestamp, err := time.Parse(time.RFC3339, ratingRequest.TimeStamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp format: %v", err)
	}

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	// Insert candidate ratings
	for _, rating := range ratingRequest.Ratings {
		_, err := tx.Exec(`
			INSERT INTO Rating (UserId, TimeStamp, RatingCardId, RatingCandidate, TextResponseCandidate, RatingEmployer, TextResponseEmployer)
			VALUES (?, ?, ?, ?, ?, NULL, NULL)`,
			ratingRequest.UserId, timestamp, rating.RatingCardId, rating.RatingCandidate, rating.TextResponseCandidate)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error saving rating: %v", err)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}
