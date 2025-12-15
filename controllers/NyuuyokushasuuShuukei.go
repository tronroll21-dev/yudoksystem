package controllers

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	controllerHelpers "tronroll21-dev/yudoksystem/controllers/helpers"
	"tronroll21-dev/yudoksystem/models"
)

func NyuuyokushasuuShuukeiHandler(c *gin.Context) {

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	if start_date == "" {
		c.String(http.StatusBadRequest, "開始日パラメータがありません。")
		return
	}

	if end_date == "" {
		c.String(http.StatusBadRequest, "最終日パラメータがありません。")
		return
	}

	NyuuyokushasuuShuukeiData, err := models.GetNyuuyokushasuuShuukeiDataByDate(start_date, end_date)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get uriage nikkei data"+err.Error())
		return
	}

	NyuuyokushasuuRuikeiData, err := models.GetNyuuyokushasuuRuikeiDataByEndDate(end_date)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get uriage nikkei data"+err.Error())
		return
	}

	// // Render the template to a buffer instead of directly to the response
	tmpl, err := controllerHelpers.ParseTemplateWithFunc("templates/nyuuyokushasuu.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to parse HTML template: %v", err)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "nyuuyokushasuu.html", gin.H{
		"NyuuyokushasuuDailyData":   NyuuyokushasuuShuukeiData.NyuuyokushasuuDailyData,
		"NyuuyokushasuuShuukeiData": NyuuyokushasuuShuukeiData,
		"NyuuyokushasuuRuikeiData":  NyuuyokushasuuRuikeiData,
	}); err != nil {
		log.Printf("Failed to execute template: %v", err)
		c.String(http.StatusInternalServerError, "Failed to render template")
		return
	}
	//	reportHTML := buf.Bytes()

	c.HTML(http.StatusOK, "nyuuyokushasuu.html", gin.H{
		"NyuuyokushasuuDailyData":   NyuuyokushasuuShuukeiData.NyuuyokushasuuDailyData,
		"NyuuyokushasuuShuukeiData": NyuuyokushasuuShuukeiData,
		"NyuuyokushasuuRuikeiData":  NyuuyokushasuuRuikeiData,
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
