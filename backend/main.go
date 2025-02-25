package main

import (
	"backend/api"
	"backend/config"
	"backend/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()
	defer db.DB()

	router := gin.Default()
	router.Use(middlewares.EnableCORS())

	api.SetupRoutes(router)

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}