package auth_usecase

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/auth"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
}

func NewAuthUsecase(
	repository domain_repository.UserRepositoryInterface,
) domain_usecase.AuthUsecaseInterface {
	return &AuthUsecase{
		UserRepository: repository,
	}
}

func (a *AuthUsecase) Execute(ctx context.Context, input dto.AuthInputDto) (domain_response.AuthResponse, error) {
	const LocalAuthProvider = "local"

	user, err := a.UserRepository.GetByEmailAndAuthProvider(ctx, input.Email, LocalAuthProvider)

	if user == nil {
		return domain_response.AuthResponse{}, errors.New("user not found")
	}

	if err != nil {
		return domain_response.AuthResponse{}, err
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(input.Password))
	if err != nil {
		return domain_response.AuthResponse{}, err
	}

	token, err := generateToken(input.Email, user.Uuid, LocalAuthProvider)

	if err != nil {
		return domain_response.AuthResponse{}, err
	}

	return domain_response.AuthResponse{Token: token}, nil
}

func generateToken(email string, userUuid string, authProvider string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_email":    email,
		"user_uuid":     userUuid,
		"auth_provider": authProvider,
		"exp":           time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(configs.ApplicationCfg.JwtSecret))
}
