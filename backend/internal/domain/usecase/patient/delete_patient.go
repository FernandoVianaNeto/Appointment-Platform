package domain_usecase_patient

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type DeletePatientUsecaseInterface interface {
	Execute(ctx context.Context, input dto.DeletePatientInputDto) error
}
