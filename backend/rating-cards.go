package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RatingCards struct {
	ID            string `json:"id" gorm:"primaryKey"`
	Question      string `json:"question"`
	Category      string `json:"category"`
	OrderId       int32   `json:"orderId"`
}

var db *gorm.DB

func init() {
    dsn := "devuser:devpassword@tcp(127.0.0.1:3306)/fox_and_hound?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&RatingCards{}); err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}
}


func getRatingCardDtoObject(w http.ResponseWriter, r *http.Request) {
	responses := getRatingCards()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func getRatingCards() []RatingCards {
    var ratingCards []RatingCards
    result := db.Find(&ratingCards)
    if result.Error != nil {
        log.Println("Error fetching rating cards:", result.Error)
        return []RatingCards{}
    }
    log.Printf("Found %d rating cards in the database", len(ratingCards))
    return ratingCards
}
