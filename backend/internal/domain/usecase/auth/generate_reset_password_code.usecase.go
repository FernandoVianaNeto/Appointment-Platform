package domain_auth_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type GenerateResetPasswordCodeUsecaseInterface interface {
	Execute(ctx context.Context, input dto.GenerateResetPasswordCodeInputDto) error
}
