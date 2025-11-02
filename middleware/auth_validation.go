package middleware

import (
	"bookstore/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || strings.HasPrefix(authHeader, "Bearer ") {
			utils.Error(c, http.StatusUnauthorized, "missing or invalid token")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateJWT(tokenString)

		if err != nil {
			utils.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			utils.Error(c, http.StatusUnauthorized, "expired token")
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
