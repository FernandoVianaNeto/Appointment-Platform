package middleware_test

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"appointment-platform-backend-backend/internal/infra/web/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	secret := "mysecret123"
	configs.InitializeConfigs()
	configs.ApplicationCfg.JwtSecret = secret

	middlewareFunc := middleware.JwtAuthMiddleware()

	handler := func(c *gin.Context) {
		email, _ := c.Get("user_email")
		uuid, _ := c.Get("user_uuid")
		c.JSON(http.StatusOK, gin.H{
			"email": email,
			"uuid":  uuid,
		})
	}

	r := gin.New()
	r.Use(middlewareFunc)
	r.GET("/", handler)

	generateToken := func(expired bool) string {
		claims := &middleware.JwtClaims{
			Email: "user@example.com",
			Uuid:  "uuid-1234",
			Role:  "admin",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			},
		}
		if expired {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(-time.Hour))
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(secret))
		return tokenString
	}

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
		expectedEmail  string
		expectedUuid   string
	}{
		{"No Authorization header", "", http.StatusUnauthorized, "", ""},
		{"Invalid token", "Bearer invalidtoken", http.StatusUnauthorized, "", ""},
		{"Valid token", "Bearer " + generateToken(false), http.StatusOK, "user@example.com", "uuid-1234"},
		{"Expired token", "Bearer " + generateToken(true), http.StatusUnauthorized, "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				assert.Contains(t, w.Body.String(), tt.expectedEmail)
				assert.Contains(t, w.Body.String(), tt.expectedUuid)
			}
		})
	}
}
