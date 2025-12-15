package controllers

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	controllerHelpers "tronroll21-dev/yudoksystem/controllers/helpers"
	"tronroll21-dev/yudoksystem/models"
)

func UriageNikkeiReportHandler(c *gin.Context) {

	// dateStr := c.Query("date")
	// if dateStr == "" {
	// 	c.String(http.StatusBadRequest, "Date parameter is missing")
	// 	return
	// }
	// date, err := time.Parse("2006-01-02", dateStr)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Invalid date format")
	// 	return
	// }

	// reportHTML, err := models.GetUriageNikkeiReportByDate(time.Date(2025, 8, 21, 0, 0, 0, 0, time.UTC), time.Date(2025, 9, 20, 0, 0, 0, 0, time.UTC))
	// if err != nil {
	// 	log.Printf("Error generating sales report: %v", err)
	// 	c.String(http.StatusInternalServerError, "Failed to generate report")
	// 	return
	// }

	UriagenikkeiData, err := models.GetUriageNikkeiDataByDate(time.Date(2025, 8, 21, 0, 0, 0, 0, time.UTC).Format("2006-01-02"), time.Date(2025, 9, 20, 0, 0, 0, 0, time.UTC).Format("2006-01-02"))
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get uriage nikkei data")
		return
	}

	UriageRuikeiData, err := models.GetUriageRuikeiDataByDate(time.Date(2025, 8, 21, 0, 0, 0, 0, time.UTC).Format("2006-01-02"), time.Date(2025, 9, 20, 0, 0, 0, 0, time.UTC).Format("2006-01-02"))
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get uriage ruikei data"+err.Error())
		return
	}

	// // Render the template to a buffer instead of directly to the response
	tmpl, err := controllerHelpers.ParseTemplateWithFunc("templates/uriagenikkei.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to parse HTML template: %v", err)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "uriagenikkei.html", gin.H{
		"UriagenikkeiData": UriagenikkeiData,
	}); err != nil {
		log.Printf("Failed to execute template: %v", err)
		c.String(http.StatusInternalServerError, "Failed to render template")
		return
	}
	//	reportHTML := buf.Bytes()

	c.HTML(http.StatusOK, "uriagenikkei.html", gin.H{
		"UriagenikkeiDailyData": UriagenikkeiData.UriageNikkeiDailyData,
		"EigyouNissuu":          UriagenikkeiData.EigyouNissuu,
		"UriageRuikeiData":      UriageRuikeiData,
	})

	// pdfBytes, err := controllerHelpers.GeneratePDFfromHTML(reportHTML)
	// if err != nil {
	// 	// handle error with c.String(...)
	// 	log.Printf("Error generating PDF: %v", err)
	// 	c.String(http.StatusInternalServerError, "Failed to generate PDF")
	// 	return

	// }

	// c.Header("Content-Type", "application/pdf")
	// c.Header("Content-Disposition", "attachment; filename=report.pdf")
	// c.Data(http.StatusOK, "application/pdf", pdfBytes)

}
