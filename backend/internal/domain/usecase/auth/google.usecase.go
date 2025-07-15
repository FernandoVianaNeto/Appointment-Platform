package domain_auth_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
)

type GoogleAuthUsecaseInterface interface {
	Execute(ctx context.Context, input dto.GoogleAuthInputDto) (domain_response.AuthResponse, error)
}
