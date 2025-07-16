package appointments

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type CreateAppointmentUsecaseInterface interface {
	Execute(ctx context.Context, input dto.CreateAppointmentInputDto) error
}
