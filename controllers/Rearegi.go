package controllers

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	controllerHelpers "tronroll21-dev/yudoksystem/controllers/helpers"
	"tronroll21-dev/yudoksystem/models"
)

// func ShiiremeisaiHandler(c *gin.Context) {

// 	year_month := c.Query("year_month") // Assuming you want to get the first day of the month

// 	shiiremeisaiData, err := models.FindAllShiiremeisai(year_month)
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Failed to get shiiremeisai data")
// 		return
// 	}

// 	c.JSON(http.StatusOK, shiiremeisaiData)

// }

type RearegiDetails struct {
	Hiduke  string                 `json:"hiduke"`
	Records []models.RearegiDetail `json:"records"`
}

type GroupedRearegiDetails struct {
	Details           []models.RearegiDetail
	GroupRank         int
	GroupTotalKingaku int
	GroupTotalSuryo   int
}

func SaveRearegiDetailHandler(c *gin.Context) {
	var records RearegiDetails
	if err := c.ShouldBindJSON(&records); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.SaveRearegiDetail(records.Records, records.Hiduke); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save rearegi detail data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully saved", "records": records})
}

func GetRearegiDetailHandler(c *gin.Context) {

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	if start_date == "" || end_date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start date and end date parameters are required"})
		return
	}
	details, err := models.FindRearegiDetails(start_date, end_date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rearegi detail data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"start_date": start_date, "end_date": end_date, "records": details})
}

// GetRearegiReportHandler handles the /rearegi-report GET request
func GetRearegiReportHandler(c *gin.Context) {

	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	details, err := models.FindRearegiDetails(start_date, end_date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rearegi detail data"})
		return
	}

	var groupedDetails []GroupedRearegiDetails
	var currentGroup *GroupedRearegiDetails
	var grandTotalKingaku int
	var grandTotalSuryo int

	for _, detail := range details {
		if currentGroup == nil || currentGroup.GroupRank != detail.Rank {
			if currentGroup != nil {
				groupedDetails = append(groupedDetails, *currentGroup)
			}
			currentGroup = &GroupedRearegiDetails{
				GroupRank: detail.Rank,
			}
		}

		currentGroup.Details = append(currentGroup.Details, detail)
		currentGroup.GroupTotalKingaku += detail.Kingaku
		currentGroup.GroupTotalSuryo += detail.Suryo
		grandTotalKingaku += detail.Kingaku
		grandTotalSuryo += detail.Suryo
	}

	if currentGroup != nil {
		groupedDetails = append(groupedDetails, *currentGroup)
	}

	var title_period string
	if start_date == end_date {
		title_period = start_date
	} else {
		title_period = start_date + "～" + end_date
	}

	title := title_period

	// Render the template to a buffer instead of directly to the response
	tmpl, err := controllerHelpers.ParseTemplateWithFunc("templates/rearegimeisai.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to parse HTML template: %v", err)
		return
	}

	// fmt.Printf("Grouped Details: %+v\n", groupedDetails)

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "rearegimeisai.html", gin.H{
		"Title":             title,
		"GroupedDetails":    groupedDetails,
		"GrandTotalKingaku": grandTotalKingaku,
		"GrandTotalSuryo":   grandTotalSuryo,
	}); err != nil {
		log.Printf("Failed to execute template: %v", err)
		c.String(http.StatusInternalServerError, "Failed to render template")
		return
	}

	c.Header("Content-Type", "text/html")
	c.Data(http.StatusOK, "text/html", buf.Bytes())

	// reportHTML := buf.Bytes()

	// pdfBytes, err := controllerHelpers.GeneratePDFfromHTML(reportHTML)
	// if err != nil {
	// 	// handle error with c.String(...)
	// 	log.Printf("Error generating PDF: %v", err)
	// 	c.String(http.StatusInternalServerError, "Failed to generate PDF")
	// 	return

	// }

	// c.Header("Content-Type", "application/pdf")
	// c.Data(http.StatusOK, "application/pdf", pdfBytes)

}
