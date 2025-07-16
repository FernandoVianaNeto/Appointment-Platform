package appointments

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type EditAppointmentUsecaseInterface interface {
	Execute(ctx context.Context, input dto.EditAppointmentInputDto) error
}
