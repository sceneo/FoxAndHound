package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/rating-cards", getRatingCardDtoObject)
	mux.HandleFunc("/api/rating-cards/save", saveRatingData) // New POST endpoint
	handler := enableCORS(mux)

	log.Println("The train has no break! Fox and hound is on its way. Ordering beers already...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
