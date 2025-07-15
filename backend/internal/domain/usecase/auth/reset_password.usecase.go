package domain_auth_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type ResetPasswordUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ResetPasswordInputDto) error
}
