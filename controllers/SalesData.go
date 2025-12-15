package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	//controllerHelpers "tronroll21-dev/yudoksystem/controllers/helpers"
	"tronroll21-dev/yudoksystem/models"
)

func SalesDataHandler(c *gin.Context) {
	//jsonBytese selected date from the query parameters
	dateStr := c.Query("date")
	log.Printf("Received HTMX request for date: %s", dateStr)

	// Parse the date string into a time.Time object
	// We'll use the YYYY-MM-DD format from the HTML //datepicker
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		// If parsing fails, return a bad request status
		c.String(http.StatusBadRequest, "Invalid date format")
		return
	}

	// Use the models package to get the sales record for the given date
	record, found, err := models.GetSalesRecordByDate(date)
	if err != nil {
		log.Printf("Error getting sales record: %v", err)
		c.String(http.StatusInternalServerError, "Database error")
		return
	}

	// Prepare the data to be returned as JSON
	data := gin.H{
		"Record": record,
		"Found":  found,
		"Mode":   "登録", // Default mode is "登録" (register)
		"Color":  "green",
	}

	if found {
		// If a record was found, update the mode and color
		data["Mode"] = "更新" // "更新" (update)
		data["Color"] = "orange"
	}

	// Return the data as JSON
	c.JSON(http.StatusOK, data)
}

func PostSalesDataHandler(c *gin.Context) {
	var input models.DailyReportRaw
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	record, found, mode, err := models.InsertOrUpdateSalesRecord(&input)
	if err != nil {
		log.Printf("Error inserting/updating sales record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	data := gin.H{
		"Record": record,
		"Found":  found,
		"Mode":   mode, // Default mode is "登録" (register)
		"Color":  "green",
	}

	if found {
		// If a record was found, update the mode and color
		data["Mode"] = "更新" // "更新" (update)
		data["Color"] = "orange"
	}

	// Return the data as JSON
	c.JSON(http.StatusOK, data)
}
