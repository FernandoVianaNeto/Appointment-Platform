package appointment_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	"context"
)

type SetAppointmentStatusUsecase struct {
	AppointmentRepository domain_repository.AppointmentRepositoryInterface
}

func NewSetAppointmentStatusUsecase(
	repository domain_repository.AppointmentRepositoryInterface,
) domain_usecase_appointment.SetAppointmentStatusUsecaseInterface {
	return &SetAppointmentStatusUsecase{
		AppointmentRepository: repository,
	}
}

func (u *SetAppointmentStatusUsecase) Execute(ctx context.Context, input dto.SetAppointmentStatusInputDto) error {
	err := u.AppointmentRepository.UpdateStatus(ctx, input.Status, input.Uuid)

	return err
}
