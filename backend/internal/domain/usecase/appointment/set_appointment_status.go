package domain_usecase_appointment

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type SetAppointmentStatusUsecaseInterface interface {
	Execute(ctx context.Context, input dto.SetAppointmentStatusInputDto) error
}
