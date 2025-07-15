package web

import (
	"appointment-platform-backend-backend/internal/infra/web/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine, server *Server) *gin.Engine {
	{
		auth := engine.Group("/auth")
		{
			auth.POST("/", server.AuthHandler)
			auth.POST("/google", server.GoogleAuthHandler)
			auth.POST("/generate-reset-code", server.GenerateResetPasswordCodeHandler)
			auth.POST("/reset-password", server.ResetPasswordHandler)
			auth.POST("/validate-code", server.ValidateResetPasswordCode)
		}
	}

	{
		user := engine.Group("/user")
		{
			user.POST("/create", server.CreateUserHandler)
			user.POST("/create/google", server.CreateGoogleUserHandler)
			user.GET("/:uuid/profile", middleware.JwtAuthMiddleware(), server.GetUserProfileHandler)
			user.PUT("/", middleware.JwtAuthMiddleware(), server.UpdateUserHandler)
		}
	}

	{
		heathCheck := engine.Group("/health")
		{
			heathCheck.GET("/check", server.HealthCheckHandler)
		}
	}

	return engine
}
