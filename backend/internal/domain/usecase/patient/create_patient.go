package domain_usecase_patient

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type CreatePatientUsecaseInterface interface {
	Execute(ctx context.Context, input dto.CreatePatientInputDto) error
}
