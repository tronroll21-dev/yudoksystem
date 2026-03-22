package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"tronroll21-dev/yudoksystem/models"

	"github.com/gin-gonic/gin"
)

func GasReadingsHandler(c *gin.Context) {
	yearStr := c.Query("year")
	monthStr := c.Query("month")
	log.Printf("Received HTMX request for year: %s, month: %s", yearStr, monthStr)

	record, found, err := models.GetGasReadingsByYearAndMonth(yearStr, monthStr)
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

func SaveGasReadingHandler(c *gin.Context) {
	var payload struct {
		ID             int    `json:"id"`
		Year           int    `json:"year"`
		Month          int    `json:"month"`
		Day            int    `json:"day"`
		Boiler1Reading int    `json:"boiler1_reading"`
		Boiler2Reading int    `json:"boiler2_reading"`
		Boiler3Reading int    `json:"boiler3_reading"`
		Boiler4Reading int    `json:"boiler4_reading"`
		TansanReading  int    `json:"tansan_reading"`
		Author         string `json:"author"`
		Memo           string `json:"memo"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	reading := models.GasReading{
		ID:             payload.ID,
		Year:           payload.Year,
		Month:          payload.Month,
		Day:            payload.Day,
		Boiler1Reading: payload.Boiler1Reading,
		Boiler2Reading: payload.Boiler2Reading,
		Boiler3Reading: payload.Boiler3Reading,
		Boiler4Reading: payload.Boiler4Reading,
		TansanReading:  payload.TansanReading,
		Author:         payload.Author,
		Memo:           sql.NullString{String: payload.Memo, Valid: payload.Memo != ""},
	}

	if err := models.SaveGasReading(&reading); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
