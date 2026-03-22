package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tronroll21-dev/yudoksystem/models"
)

func ShiiremeisaiHandler(c *gin.Context) {

	year_month := c.Query("year_month") // Assuming you want to get the first day of the month

	shiiremeisaiData, err := models.FindAllShiiremeisai(year_month)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get shiiremeisai data")
		return
	}

	c.JSON(http.StatusOK, shiiremeisaiData)

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
