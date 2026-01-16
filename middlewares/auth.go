package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"tronroll21-dev/yudoksystem/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		// Find the user
		var user *models.User
		userID := uint(claims["sub"].(float64))
		user, err := models.GetUserById(userID)

		if err != nil || user.ID == 0 {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		// Attach to req
		c.Set("user", *user)

		// Continue
		c.Next()
	} else {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
	}
}
