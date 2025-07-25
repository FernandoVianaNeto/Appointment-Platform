package appointment_usecase

import (
	adapter "appointment-platform-backend-backend/internal/domain/adapters/email_sender"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	"context"
)

type GetNextAppointmentsAndSendReminder struct {
	AppointmentRepository domain_repository.AppointmentRepositoryInterface
	EmailSender           adapter.EmailSenderAdapterInterface
}

func NewGetNextAppointmentsAndSendReminder(
	repository domain_repository.AppointmentRepositoryInterface,
	emailSender adapter.EmailSenderAdapterInterface,
) domain_usecase_appointment.GetNextAppointmentsAndSendReminder {
	return &GetNextAppointmentsAndSendReminder{
		AppointmentRepository: repository,
		EmailSender:           emailSender,
	}
}

func (u *GetNextAppointmentsAndSendReminder) Execute(ctx context.Context) domain_response.GetNextAppointmentsResponse {
	return domain_response.GetNextAppointmentsResponse{}
}
