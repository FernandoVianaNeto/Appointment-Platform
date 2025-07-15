package domain_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
)

type GetUserProfileUsecaseInterface interface {
	Execute(ctx context.Context, input dto.GetUserInputDto) (*domain_response.GetUserProfileResponse, error)
}
