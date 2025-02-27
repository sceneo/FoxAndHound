package main

import (
	"backend/api"
	"backend/config"
	"backend/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	_ "backend/docs"
)

// @title Fox & Hound API
// @version 1.0
// @description This is a Gin-based API for the Fox & Hound application.
// @host localhost:8080
// @BasePath /api
func main() {
	db := config.ConnectDatabase()
	
	dbInstance, err := db.DB()
	if err != nil {
		log.Fatal("❌ Failed to get database instance:", err)
	}
	defer dbInstance.Close()

	router := gin.Default()
	router.Use(middlewares.EnableCORS())

	api.SetupRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("🚀 Server running on http://localhost:8080")
	log.Println("📖 Swagger UI available at http://localhost:8080/swagger/index.html")
	log.Fatal(router.Run(":8080"))
}