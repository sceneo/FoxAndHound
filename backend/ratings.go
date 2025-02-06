package main

import (
	"database/sql"
	"fmt"
)

// TotalRating holds the structure of total ratings information.
type TotalRating struct {
	RatingCardId     int     `json:"ratingCardId"`
	AverageRating    float64 `json:"averageRating"`
	TotalSubmissions int     `json:"totalSubmissions"`
}

// SaveTotalRating calculates and saves the total rating for each rating card.
func SaveTotalRating(db *sql.DB) error {
	// Get all the ratings from the database
	rows, err := db.Query(`
		SELECT RatingCardId, AVG(RatingCandidate) AS AverageRating, COUNT(*) AS TotalSubmissions
		FROM Rating
		GROUP BY RatingCardId`)
	if err != nil {
		return fmt.Errorf("error fetching ratings: %v", err)
	}
	defer rows.Close()

	// Insert or update the total ratings
	for rows.Next() {
		var totalRating TotalRating
		if err := rows.Scan(&totalRating.RatingCardId, &totalRating.AverageRating, &totalRating.TotalSubmissions); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}

		// Insert total rating into the database (if not already present)
		_, err := db.Exec(`
			INSERT INTO TotalRatings (RatingCardId, AverageRating, TotalSubmissions)
			VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE AverageRating = ?, TotalSubmissions = ?`,
			totalRating.RatingCardId, totalRating.AverageRating, totalRating.TotalSubmissions,
			totalRating.AverageRating, totalRating.TotalSubmissions)
		if err != nil {
			return fmt.Errorf("error saving total rating: %v", err)
		}
	}

	return nil
}
