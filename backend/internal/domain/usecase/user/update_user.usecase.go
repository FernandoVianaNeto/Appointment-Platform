package domain_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type UpdateUserUsecaseInterface interface {
	Execute(ctx context.Context, input dto.UpdateUserInputDto) error
}
