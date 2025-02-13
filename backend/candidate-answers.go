package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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

func GetAllRatingRequests() ([]RatingRequest, error) {
	var db = GetDb()
	var ratings []struct {
		UserId                int
		TimeStamp             time.Time
		RatingCardId          int
		RatingCandidate       int
		TextResponseCandidate string
	}

	result := db.Raw(`
        SELECT UserId, TimeStamp, RatingCardId, RatingCandidate, TextResponseCandidate
        FROM Rating
    `).Scan(&ratings)

	if result.Error != nil {
		return nil, fmt.Errorf("error querying ratings: %v", result.Error)
	}

	ratingRequestsMap := make(map[int]RatingRequest)

	for _, rating := range ratings {
		if _, exists := ratingRequestsMap[rating.UserId]; !exists {
			ratingRequestsMap[rating.UserId] = RatingRequest{
				UserId:    rating.UserId,
				TimeStamp: rating.TimeStamp.Format(time.RFC3339),
				Ratings:   []Rating{{RatingCardId: rating.RatingCardId, RatingCandidate: rating.RatingCandidate, TextResponseCandidate: rating.TextResponseCandidate}},
			}
		} else {
			request := ratingRequestsMap[rating.UserId]
			request.Ratings = append(request.Ratings, Rating{RatingCardId: rating.RatingCardId, RatingCandidate: rating.RatingCandidate, TextResponseCandidate: rating.TextResponseCandidate})
			ratingRequestsMap[rating.UserId] = request
		}
	}

	var ratingRequests []RatingRequest
	for _, request := range ratingRequestsMap {
		ratingRequests = append(ratingRequests, request)
	}

	return ratingRequests, nil
}

func getCandidateAnswersDtoObject(w http.ResponseWriter, r *http.Request) {
	// TODO: read request and call corrct method
	responses := getCandidateAnswers("TEST")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func getCandidateAnswers(id interface{}) func(http.ResponseWriter, *http.Request) {
	return nil
}
