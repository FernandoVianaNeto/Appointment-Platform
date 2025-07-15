package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HealthCheckHandler(ctx *gin.Context) {
	fmt.Println("Health check endpoint hit")
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is running",
	})
}
