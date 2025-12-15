package controllers

import (
	"encoding/json"
	"net/http"

	"tronroll21-dev/yudoksystem/models"

	"github.com/gin-gonic/gin"
)

func TantoushasHandler(c *gin.Context) {
	tantoushas, err := models.GetTantoushas()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DB問合せに失敗しました。"})
		return
	}
	jsonData, err := json.Marshal(tantoushas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なJSONです"})
		return
	}
	c.JSON(http.StatusOK, string(jsonData))
}
