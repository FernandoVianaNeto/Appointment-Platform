package appointment_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	"context"
)

type EditAppointmentUsecase struct {
	AppointmentRepository domain_repository.AppointmentRepositoryInterface
}

func NewEditAppointmentUseCase(
	repository domain_repository.AppointmentRepositoryInterface,
) domain_usecase_appointment.EditAppointmentUsecaseInterface {
	return &EditAppointmentUsecase{
		AppointmentRepository: repository,
	}
}

func (u *EditAppointmentUsecase) Execute(ctx context.Context, input dto.EditAppointmentInputDto) error {
	err := u.AppointmentRepository.Edit(ctx, input)

	return err
}
