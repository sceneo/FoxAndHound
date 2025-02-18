package main

import (
	"encoding/json"
	"net/http"
)

// TODO: find out if we need more info about the user here
type Candidate struct {
	UserId int `json:"userId"`
}

func getAllCandidatesDtoObject(w http.ResponseWriter, r *http.Request) {
	responses := getAllCandidates()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func getAllCandidates() []Candidate {

	var ratingRequests, err = GetAllRatingRequests()

	if err != nil {
		return nil
	}

	var candidates []Candidate
	for _, request := range ratingRequests {
		candidates = append(candidates, Candidate{UserId: request.UserId})
	}

	return candidates
}
