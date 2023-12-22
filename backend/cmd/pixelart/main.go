package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Koenigseder/pixelart/internal/canvas"
	"github.com/Koenigseder/pixelart/internal/rest"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Setup Gin router
func setupRouter() *gin.Engine {
	// Define Gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	router.StaticFS("/web", http.Dir("../frontend"))

	apiGroup := router.Group("/api")
	apiGroup.GET("/pixels", rest.GetPixels) // Get all pixels
	apiGroup.POST("/pixel", rest.SetPixel)  // Set a specific pixel

	return router
}

func main() {
	router := setupRouter() // Setup router

	// Read canvas height and width from env vars
	canvasWidth, err := strconv.Atoi(os.Getenv("CANVAS_WIDTH"))
	if err != nil {
		log.Fatalf("'CANVAS_WIDTH' is not an integer!")
	}

	canvasHeight, err := strconv.Atoi(os.Getenv("CANVAS_HEIGHT"))
	if err != nil {
		log.Fatalf("'CANVAS_HEIGHT' is not an integer!")
	}

	// Check if the latest backup file should be used
	if os.Getenv("USE_LATEST_BACKUP_FILE") == "true" {
		if err := canvas.Canvas.InitializeFromLatestBackup(canvasWidth, canvasHeight); err != nil {
			log.Fatalf("Error initializing from backup file: %v\n", err)
		}
	} else {
		canvas.Canvas.InitializeEmpty(canvasWidth, canvasHeight) // Initialize empty canvas
	}

	go canvas.SaveCanvasAsFile() // Start backup routine

	// Start the Gin router
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
}
