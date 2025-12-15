package controllers

import (
	"encoding/csv"
	"log"
	"net/http"

	"tronroll21-dev/yudoksystem/models"

	"github.com/gin-gonic/gin"
)

func JissekiyosokuHandler(c *gin.Context) {
	// Call the model function to get the data
	data, err := models.GetJissekiyosokuData()
	if err != nil {
		log.Printf("Error fetching jissekiyosoku data: %v", err)
		c.String(http.StatusInternalServerError, "Failed to fetch jissekiyosoku data")
		return
	}

	// Set headers for CSV download
	c.Header("Content-Type", "text/csv; charset=utf-8")
	//

	// Create a CSV writer
	writer := csv.NewWriter(c.Writer)

	// Write all rows (including headers if present)
	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			log.Printf("Error writing CSV row: %v", err)
			break
		}
	}

	// Flush any remaining data
	writer.Flush()

	c.Status(http.StatusOK)
}
