package main

import (
	"backend/api"
	"backend/config"
	"backend/middlewares"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

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

	// Add a simple health check endpoint
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
        })
    })

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }
	log.Printf("🚀 Server running on http://localhost:%s", port)
	log.Printf("📖 Swagger UI available at http://localhost:%s/swagger/index.html", port)
	log.Fatal(router.Run(":" + port))
}