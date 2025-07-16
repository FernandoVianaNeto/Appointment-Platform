package domain_usecase_appointment

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type ListeAppointmentsUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ListAppointmentInputDto) error
}
