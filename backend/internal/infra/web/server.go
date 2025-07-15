package web

import (
	domain_auth_usecase "appointment-platform-backend-backend/internal/domain/usecase/auth"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/user"
	"context"

	gin "github.com/gin-gonic/gin"
)

type Server struct {
	router            *gin.Engine
	CreateUserUsecase domain_usecase.CreateUserUsecaseInterface
	GetUserUsecase    domain_usecase.GetUserProfileUsecaseInterface
	UpdateUserUsecase domain_usecase.UpdateUserUsecaseInterface
	AuthUseCase       domain_auth_usecase.AuthUsecaseInterface
	GoogleAuthUsecase domain_auth_usecase.GoogleAuthUsecaseInterface

	GenerateResetPasswordCodeUsecase domain_auth_usecase.GenerateResetPasswordCodeUsecaseInterface
	ResetPasswordUsecase             domain_auth_usecase.ResetPasswordUsecaseInterface
	ValidateResetPasswordCodeUsecase domain_auth_usecase.ValidateResetPasswordCodeUsecaseInterface
}

func NewServer(
	ctx context.Context,
	createUserUsecase domain_usecase.CreateUserUsecaseInterface,
	getUserUsecase domain_usecase.GetUserProfileUsecaseInterface,
	updateUserUsecase domain_usecase.UpdateUserUsecaseInterface,
	authUsecase domain_auth_usecase.AuthUsecaseInterface,
	googleAuthUsecase domain_auth_usecase.GoogleAuthUsecaseInterface,
	generateResetPasswordCodeUsecase domain_auth_usecase.GenerateResetPasswordCodeUsecaseInterface,
	resetPasswordUsecase domain_auth_usecase.ResetPasswordUsecaseInterface,
	validateResetPasswordCodeUsecase domain_auth_usecase.ValidateResetPasswordCodeUsecaseInterface,
) *Server {
	router := gin.Default()

	server := &Server{
		CreateUserUsecase:                createUserUsecase,
		GetUserUsecase:                   getUserUsecase,
		UpdateUserUsecase:                updateUserUsecase,
		AuthUseCase:                      authUsecase,
		GoogleAuthUsecase:                googleAuthUsecase,
		GenerateResetPasswordCodeUsecase: generateResetPasswordCodeUsecase,
		ResetPasswordUsecase:             resetPasswordUsecase,
		ValidateResetPasswordCodeUsecase: validateResetPasswordCodeUsecase,
	}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
