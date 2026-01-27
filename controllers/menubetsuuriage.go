package controllers

import (
	"bytes"
	"log"
	"net/http"

	controllerHelpers "tronroll21-dev/yudoksystem/controllers/helpers"
	"tronroll21-dev/yudoksystem/models"

	"github.com/gin-gonic/gin"
)

// MenubetsuUriageHandler handles the /menubetsuuriage GET request
func MenubetsuUriageHandler(c *gin.Context) {

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")
	groupedSale, err := models.GetMenubetsuUriage(start_date, end_date)
	if err != nil {
		log.Printf("Error generating sales report: %v", err)
		c.String(http.StatusInternalServerError, "Failed to generate report")
		return
	}

	var title_period string
	if start_date == end_date {
		title_period = start_date
	} else {
		title_period = start_date + "～" + end_date
	}

	title := "メニュー別売上　" + title_period

	// Render the template to a buffer instead of directly to the response
	tmpl, err := controllerHelpers.ParseTemplateWithFunc("templates/menubetsuuriage.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to parse HTML template: %v", err)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "menubetsuuriage.html", gin.H{
		"Title":       title,
		"GroupedSale": groupedSale,
	}); err != nil {
		log.Printf("Failed to execute template: %v", err)
		c.String(http.StatusInternalServerError, "Failed to render template")
		return
	}
	reportHTML := buf.Bytes()

	pdfBytes, err := controllerHelpers.GeneratePDFfromHTML(reportHTML)
	if err != nil {
		// handle error with c.String(...)
		log.Printf("Error generating PDF: %v", err)
		c.String(http.StatusInternalServerError, "Failed to generate PDF")
		return

	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=report.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)

}

func SaveMenubetsuUriage(c *gin.Context) {
	var reqBody struct {
		Date         string              `json:"date"`
		SoldProducts []models.SoldProduct `json:"soldProducts"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := models.SaveMenubetsuUriage(reqBody.Date, reqBody.SoldProducts); err != nil {
		log.Printf("Failed to save menubetsu uriage: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
