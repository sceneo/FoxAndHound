package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// saveRatingData handles the POST request to save rating data
func saveRatingData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var ratingRequest RatingRequest
	err := json.NewDecoder(r.Body).Decode(&ratingRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request: %v", err), http.StatusBadRequest)
		return
	}

	// Database connection (replace with actual credentials)
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ratingdb")
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Save candidate ratings
	err = SaveCandidateRatings(db, ratingRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving ratings: %v", err), http.StatusInternalServerError)
		return
	}

	// Save total ratings
	err = SaveTotalRating(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving total ratings: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ratings saved successfully"))
}
