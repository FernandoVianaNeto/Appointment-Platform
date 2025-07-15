package domain_auth_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
)

type AuthUsecaseInterface interface {
	Execute(ctx context.Context, input dto.AuthInputDto) (domain_response.AuthResponse, error)
}
