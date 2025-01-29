package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CategoryEnum string

const (
	CategoryFood     CategoryEnum = "Food"
	CategoryElectronics CategoryEnum = "Electronics"
	CategoryClothing CategoryEnum = "Clothing"
)

type ResponseObject struct {
	Name          string       `json:"name"`
	Category      CategoryEnum `json:"category"`
	AverageRating float64      `json:"averageRating"`
}

func getObjectHandler(w http.ResponseWriter, r *http.Request) {
	response := ResponseObject{
		Name:          "Sample Product",
		Category:      CategoryFood,
		AverageRating: 4.5,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/object", getObjectHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
