package domain_usecase_appointment

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type DeleteAppointmentUsecaseInterface interface {
	Execute(ctx context.Context, input dto.DeleteAppointmentInputDto) error
}
