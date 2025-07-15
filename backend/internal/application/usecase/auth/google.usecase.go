package auth_usecase

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/auth"
	"context"
	"errors"
	"fmt"

	"google.golang.org/api/idtoken"
)

type GoogleAuthUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
}

func NewGoogleAuthUsecase(
	repository domain_repository.UserRepositoryInterface,
) domain_usecase.GoogleAuthUsecaseInterface {
	return &GoogleAuthUsecase{
		UserRepository: repository,
	}
}

func (a *GoogleAuthUsecase) Execute(ctx context.Context, input dto.GoogleAuthInputDto) (domain_response.AuthResponse, error) {
	payload, err := idtoken.Validate(ctx, input.Token, configs.GoogleAuthCfg.ClientId)

	if err != nil {
		return domain_response.AuthResponse{}, fmt.Errorf("invalid token: %v", err)
	}

	emailRaw, ok := payload.Claims["email"]
	if !ok {
		return domain_response.AuthResponse{}, fmt.Errorf("email not found in token claims")
	}

	email, ok := emailRaw.(string)
	if !ok {
		return domain_response.AuthResponse{}, fmt.Errorf("email is not a valid string")
	}

	user, err := a.UserRepository.GetByEmailAndAuthProvider(ctx, email, "google")

	if user.Uuid == "" {
		return domain_response.AuthResponse{}, errors.New("user not found")
	}

	if err != nil {
		return domain_response.AuthResponse{}, err
	}

	token, err := generateToken(email, payload.Issuer, "google")

	if err != nil {
		return domain_response.AuthResponse{}, err
	}

	return domain_response.AuthResponse{Token: token}, nil
}

func VerifyGoogleIDToken(ctx context.Context, idToken string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(ctx, idToken, configs.GoogleAuthCfg.ClientId)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
