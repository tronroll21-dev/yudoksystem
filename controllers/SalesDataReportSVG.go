package controllers

import (
	"bytes"
	"fmt"
	"log"
	"net/http" // Use text/template for SVG
	"time"

	controllerHelpers "tronroll21-dev/yudoksystem/controllers/helpers"
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
	tmpl, err := controllerHelpers.ParseSVGTemplateWithFunc("templates/uriagenichijihoukokusho.svg")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		c.String(http.StatusInternalServerError, "Failed to parse template")
		return
	}

	// Execute the template into a buffer
	var buf bytes.Buffer
	now := time.Now()

	err = tmpl.ExecuteTemplate(&buf, "uriagenichijihoukokusho.svg", gin.H{
		"ReportData": ReportData,
		"Timestamp": fmt.Sprintf("%d年%d月%d日　　%d:%02d:%02d\n",
			now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second()),
	})
	if err != nil {
		log.Printf("Error executing template: %v", err)
		c.String(http.StatusInternalServerError, "Failed to execute template")
		return
	}

	reportSVG := buf.Bytes()

	pdfBytes, err := controllerHelpers.GeneratePDFfromSVG(reportSVG)
	if err != nil {
		// handle error with c.String(...)
		log.Printf("Error generating PDF: %v", err)
		c.String(http.StatusInternalServerError, "Failed to generate PDF")
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)

	// Set the content type and return the SVG data
	//c.Header("Content-Type", "image/svg+xml")
	//c.Header("Content-Type", "application/json")
	//c.Data(http.StatusOK, "image/svg+xml", buf.Bytes())
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": ReportData,
	// })
}
