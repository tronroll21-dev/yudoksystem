package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"tronroll21-dev/yudoksystem/models"
)

// CreateRange は新しい料金範囲データを作成します (POST /api/ranges)
func CreateRange(c *gin.Context) {
	var newRange models.PriceRange
	// リクエストボディからJSONをバインド
	if err := c.ShouldBindJSON(&newRange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストボディの形式が不正です: " + err.Error()})
		return
	}

	// モデル層に処理を委譲
	createdRange, err := models.InsertRange(newRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データの挿入に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdRange)
}

// GetRanges は全ての料金範囲データを取得します (GET /api/ranges)
func GetRanges(c *gin.Context) {
	// モデル層に処理を委譲
	ranges, err := models.FindAllRanges()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データの取得に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, ranges)
}

// GetRangeByID は指定されたIDの料金範囲データを取得します (GET /api/ranges/:id)
func GetRangeByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID形式です"})
		return
	}

	// モデル層に処理を委譲
	r, err := models.FindRangeByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "指定されたIDのデータが見つかりません"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データの取得に失敗しました: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, r)
}

// UpdateRange は指定されたIDの料金範囲データを更新します (PUT /api/ranges/:id)
func UpdateRange(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID形式です"})
		return
	}

	var updatedRange models.PriceRange
	// リクエストボディからJSONをバインド
	if err := c.ShouldBindJSON(&updatedRange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストボディの形式が不正です: " + err.Error()})
		return
	}

	// モデル層に処理を委譲
	resultRange, err := models.UpdateRange(id, updatedRange)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "更新対象のデータが見つからないか、データが変更されていません"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データの更新に失敗しました: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, resultRange)
}

// DeleteRange は指定されたIDの料金範囲データを削除します (DELETE /api/ranges/:id)
func DeleteRange(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID形式です"})
		return
	}

	// モデル層に処理を委譲
	if err := models.DeleteRange(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "削除対象のデータが見つかりません"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データの削除に失敗しました: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "データを正常に削除しました", "id": id})
}
