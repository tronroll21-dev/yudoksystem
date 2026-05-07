// main.go
package main

import (
	"html/template"
	"log"
	"tronroll21-dev/yudoksystem/models"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	// Assuming the controllers package is in the same directory

	// Assuming the models package is in the same directory

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func formatNumber(n int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", n)
}

// Main function to set up the Gin router and start the server
func main() {

	err := godotenv.Load()
	if err != nil {
		// Log a fatal error if the .env file cannot be loaded, as the app
		// cannot run without its configuration.
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := models.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create a new Gin router with the default middleware
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"format": formatNumber,
	})

	registerRoutes(router)

	log.Println("Starting server on :8080")
	router.Run(":8080")
}
