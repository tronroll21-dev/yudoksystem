package controllers

import (
	"net/http"
	"strconv"

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

}

func SaveShiiremeisaiHandler(c *gin.Context) {
	var record models.Shiiremeisai
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.SaveShiiremeisai(record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save shiiremeisai data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully saved", "record": record})
}

func DeleteShiiremeisaiHandler(c *gin.Context) {
	var record models.Shiiremeisai
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DeleteShiiremeisai(record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete shiiremeisai data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted", "record": record})
}

func SaveShiireshouhizeiHandler(c *gin.Context) {
	var body struct {
		YearMonth    int    `json:"year_month"`
		ContractorID int    `json:"contractor_id"`
		Shouhizei    string `json:"shouhizei"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.SaveShiireshouhizei(body.YearMonth, body.ContractorID, body.Shouhizei); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save shouhizei data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully saved shouhizei"})
}

func GetAvailableContractorsHandler(c *gin.Context) {
	yearMonthStr := c.Query("year_month")
	yearMonth, _ := strconv.Atoi(yearMonthStr)

	contractors, err := models.GetAvailableContractors(yearMonth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch available contractors"})
		return
	}
	c.JSON(http.StatusOK, contractors)
}

func AddContractorHandler(c *gin.Context) {
	var body struct {
		YearMonth    int    `json:"year_month"`
		ContractorID int    `json:"contractor_id"`
		Tekiyou      string `json:"tekiyou"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.AddContractorToMonth(body.YearMonth, body.ContractorID, body.Tekiyou); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add contractor"})
		return
	}

	// Return updated records for the month
	records, err := models.FindAllShiiremeisai(strconv.Itoa(body.YearMonth))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated data"})
		return
	}
	c.JSON(http.StatusOK, records)
}

func InitializeShiiremeisaiHandler(c *gin.Context) {
	var body struct {
		YearMonth int `json:"year_month"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.InitializeShiiremeisai(body.YearMonth); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize month data"})
		return
	}

	// Return the newly initialized records
	records, err := models.FindAllShiiremeisai(strconv.Itoa(body.YearMonth))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch initialized data"})
		return
	}

	c.JSON(http.StatusOK, records)
}
