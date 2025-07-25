package domain_usecase_appointment

import (
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
)

type GetNextAppointmentsAndSendReminder interface {
	Execute(ctx context.Context) domain_response.GetNextAppointmentsResponse
}
