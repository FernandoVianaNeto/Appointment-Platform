package domain_usecase_patient

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type EditPatientUsecaseInterface interface {
	Execute(ctx context.Context, input dto.EditPatientInputDto) error
}
