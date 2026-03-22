package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"tronroll21-dev/yudoksystem/models"
)

func PowerReadingsHandler(c *gin.Context) {
	yearStr := c.Query("year")
	monthStr := c.Query("month")
	log.Printf("Received HTMX request for year: %s, month: %s", yearStr, monthStr)

	record, found, err := models.GetPowerReadingsByYearAndMonth(yearStr, monthStr)
	if err != nil {
		log.Printf("Error getting power readings: %v", err)
		c.String(http.StatusInternalServerError, "Database error")
		return
	}

	data := gin.H{
		"Record": record,
		"Found":  found,
		"Mode":   "登録",
		"Color":  "green",
	}

	if found {
		data["Mode"] = "更新"
		data["Color"] = "orange"
	}

	c.JSON(http.StatusOK, data)
}

func SavePowerReadingHandler(c *gin.Context) {
	var payload struct {
		Year         int     `json:"year"`
		Month        int     `json:"month"`
		Day          int     `json:"day"`
		PowerReading float64 `json:"power_reading"`
		Author       string  `json:"author"`
		Memo         string  `json:"memo"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	reading := models.PowerReading{
		Year:         payload.Year,
		Month:        payload.Month,
		Day:          payload.Day,
		PowerReading: payload.PowerReading,
		Author:       payload.Author,
		Memo:         sql.NullString{String: payload.Memo, Valid: payload.Memo != ""},
	}

	if err := models.SavePowerReading(&reading); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
