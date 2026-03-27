package controllers

import (
	"bytes"
	"log"
	"net/http"
	"text/template" // Use text/template for SVG
	"time"

	"tronroll21-dev/yudoksystem/models"

	"github.com/gin-gonic/gin"
)

func SalesDataReportSVGHandler(c *gin.Context) {
	dateStr := c.Query("date")
	if dateStr == "" {
		c.String(http.StatusBadRequest, "Date parameter is missing")
		return
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid date format")
		return
	}
	ReportData, err := models.GetSalesReportByDate(date)
	if err != nil {
		log.Printf("Error fetching sales report data: %v", err)
		c.String(http.StatusInternalServerError, "Failed to fetch report data")
		return
	}

	// Parse the SVG template
	tmpl, err := template.ParseFiles("templates/uriagenichijihoukokusho.svg")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		c.String(http.StatusInternalServerError, "Failed to parse template")
		return
	}

	// Execute the template into a buffer
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, gin.H{
		"ReportData": ReportData,
	})
	if err != nil {
		log.Printf("Error executing template: %v", err)
		c.String(http.StatusInternalServerError, "Failed to execute template")
		return
	}

	// Set the content type and return the SVG data
	c.Header("Content-Type", "image/svg+xml")
	c.Data(http.StatusOK, "image/svg+xml", buf.Bytes())
}
