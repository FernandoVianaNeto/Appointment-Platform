package domain_usecase_appointment

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
)

type ListAppointmentsUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ListAppointmentInputDto) (domain_response.ListAppointmentsResponse, error)
}
