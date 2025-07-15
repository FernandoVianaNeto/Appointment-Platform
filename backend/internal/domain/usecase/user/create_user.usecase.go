package domain_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type CreateUserUsecaseInterface interface {
	Execute(ctx context.Context, input dto.CreateUserInputDto) error
}
