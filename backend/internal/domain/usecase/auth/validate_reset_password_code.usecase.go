package domain_auth_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
)

type ValidateResetPasswordCodeUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ValidateResetPasswordCodeInputDto) (domain_response.ValidateResetPasswordCodeResponse, error)
}
