package middleware

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Email string `json:"user_email"`
	Role  string `json:"user_role"`
	Uuid  string `json:"user_uuid"`
	jwt.RegisteredClaims
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := []byte(configs.ApplicationCfg.JwtSecret)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token not provided"})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected assign method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
			c.Set("user_email", claims.Email)
			c.Set("user_uuid", claims.Uuid)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		}
	}
}
