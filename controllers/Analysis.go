package controllers

import (
	"log"
	"net/http"
	"time"

	"tronroll21-dev/yudoksystem/models"

	"github.com/gin-gonic/gin"
)

func AnalysisDataHandler(c *gin.Context) {
	startStr := c.Query("start_date")
	endStr := c.Query("end_date")

	if startStr == "" {
		startStr = "2026-04-21"
	}
	if endStr == "" {
		endStr = "2026-05-20"
	}

	const layout = "2006-01-02"
	start, err := time.Parse(layout, startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD"})
		return
	}
	end, err := time.Parse(layout, endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
		return
	}

	data, err := models.GetAnalysisDataRange(start, end)
	if err != nil {
		log.Printf("Error getting analysis data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, data)
}
