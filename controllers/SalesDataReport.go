package controllers

import (
	"log"
	"net/http"
	"time"

	controllerHelpers "tronroll21-dev/yudoksystem/controllers/helpers"
	"tronroll21-dev/yudoksystem/models"

	"github.com/gin-gonic/gin"
)

func SalesDataReportHandler(c *gin.Context) {

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

	c.HTML(http.StatusOK, "report.html", gin.H{
		"ReportData": ReportData,
	})
}

func SalesDataReportHandlerPDF(c *gin.Context) {

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

	reportHTML := controllerHelpers.GenerateReportHTML(ReportData)
	pdfBytes, err := controllerHelpers.GeneratePDFfromHTML(reportHTML)
	if err != nil {
		// handle error with c.String(...)
		log.Printf("Error generating PDF: %v", err)
		c.String(http.StatusInternalServerError, "Failed to generate PDF")
		return

	}

	c.Header("Content-Type", "application/pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)

}
