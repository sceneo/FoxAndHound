package main

import (
	"encoding/json"
	"net/http"
)

type RatingCards struct {
    ID            string       `json:"id"`
	Question      string       `json:"question"`
	Category      CategoryEnum `json:"category"`
	AverageRating float64      `json:"averageRating"`
	OrderId       int32        `json:"orderId"`
}

func getRatingCardDtoObject(w http.ResponseWriter, r *http.Request) {
	responses := getRatingCards()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func getRatingCards() []RatingCards {
	return []RatingCards{
		{
		    ID: "1",
			Question:      "Would a customer see you as a senior?",
			Category:      CategoryPerformance,
			AverageRating: 4.5,
			OrderId: 1,
		},
        {
            ID: "2",
            Question:      "How do you rate your proficiency relate to your role description?",
            Category:      CategoryTechnicalSkillset,
            AverageRating: 2.5,
            OrderId: 6,
        },
	}
}
