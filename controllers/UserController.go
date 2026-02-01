package controllers

import (
	"net/http"
	"os"
	"time"

	"tronroll21-dev/yudoksystem/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		UserID   uint
		Username string
		Password string
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "パラメータが存在しません。" + err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "パスワードのハッシュ化に失敗しました。",
		})
		return
	}

	user := models.User{ID: body.UserID, Username: body.Username, Password: string(hash)}

	result, err := models.UpdatePassword(&user)

	if err != nil {
		c.JSON(500, gin.H{
			"error":  "ユーザーの作成に失敗しました。",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "ユーザーの作成に成功しました。",
		"Id":       result.ID,
		"Username": result.Username,
		"Password": result.Password,
	})

}

func Login(c *gin.Context) {

	var body struct {
		UserID   string
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "パラメータが存在しません。",
		})
		return
	}

	var user *models.User

	user, err := models.GetUserByName(body.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ユーザーが存在しません。",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "パスワードが間違っています。",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Username,
		// 有効期限を1日に設定
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "トークンの作成に失敗しました。",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	// クッキーにトークンをセット(キー, 値, 有効期限, パス, ドメイン, https, httponly)
	c.SetCookie("Authorization", tokenString, 3600*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "ログイン済み:" + user.(models.User).Username,
	})

}

func GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		// This case should ideally not be reached if RequireAuth middleware is used
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	userModel, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assert user type from context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": userModel.Username,
	})
}
