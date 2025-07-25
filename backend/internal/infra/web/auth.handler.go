package web

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	mongo_exception "appointment-platform-backend-backend/internal/infra/exceptions"
	"appointment-platform-backend-backend/internal/infra/web/requests"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AuthHandler(ctx *gin.Context) {
	var req requests.AuthRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	response, err := s.AuthUseCase.Execute(ctx, dto.AuthInputDto{Email: req.Email, Password: req.Password})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *Server) GoogleAuthHandler(ctx *gin.Context) {
	var req requests.GoogleAuthRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	response, err := s.GoogleAuthUsecase.Execute(ctx, dto.GoogleAuthInputDto{Token: req.Token})

	if err != nil {
		if err.Error() == mongo_exception.MongoNotFoundException {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found"})
			return
		}

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *Server) GenerateResetPasswordCodeHandler(ctx *gin.Context) {
	var req requests.GenerateResetPasswordCodeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := s.GenerateResetPasswordCodeUsecase.Execute(ctx, dto.GenerateResetPasswordCodeInputDto{Email: req.Email})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) ResetPasswordHandler(ctx *gin.Context) {
	var req requests.ResetPasswordRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println("request", req)
	err := s.ResetPasswordUsecase.Execute(ctx, dto.ResetPasswordInputDto{
		Code:        req.Code,
		NewPassword: req.NewPassword,
		Email:       req.Email,
	})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return

	}

	ctx.Status(http.StatusOK)
}

func (s *Server) ValidateResetPasswordCode(ctx *gin.Context) {
	var req requests.ValidateResetPasswordCodeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	response, err := s.ValidateResetPasswordCodeUsecase.Execute(ctx, dto.ValidateResetPasswordCodeInputDto{
		Code:  req.Code,
		Email: req.Email,
	})

	if err != nil {
		if err.Error() == mongo_exception.MongoNotFoundException {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "User not found"})
			return
		}

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
