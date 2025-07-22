package web

import (
	"appointment-platform-backend-backend/internal/infra/web/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine, server *Server) *gin.Engine {
	{
		auth := engine.Group("/auth")
		{
			auth.POST("/login", server.AuthHandler)
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
		patient := engine.Group("/patient", middleware.JwtAuthMiddleware())
		{
			patient.POST("/create", server.CreatePatientHandler)
			patient.GET("/list", server.ListPatientHandler)
			patient.PUT("/:uuid", server.EditPatientHandler)
			patient.DELETE("/:uuid", server.DeletePatientHandler)
		}
	}

	{
		appointment := engine.Group("/appointment", middleware.JwtAuthMiddleware())
		{
			appointment.POST("/create", server.CreateAppointmentHandler)
			appointment.GET("/list", server.ListAppointmentsHandler)
			appointment.PUT("/:uuid", server.EditAppointmentHandler)
			appointment.DELETE("/", server.DeleteAppointmentHandler)
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
