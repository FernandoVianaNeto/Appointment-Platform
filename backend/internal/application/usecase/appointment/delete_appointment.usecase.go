package appointment_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	"context"
	"fmt"
)

type DeleteAppointmentUsecase struct {
	AppointmentRepository domain_repository.AppointmentRepositoryInterface
}

func NewDeleteAppointmentUseCase(
	repository domain_repository.AppointmentRepositoryInterface,
) domain_usecase_appointment.DeleteAppointmentUsecaseInterface {
	return &DeleteAppointmentUsecase{
		AppointmentRepository: repository,
	}
}

func (u *DeleteAppointmentUsecase) Execute(ctx context.Context, input dto.DeleteAppointmentInputDto) error {
	fmt.Println("INPUT", input)

	err := u.AppointmentRepository.DeleteMany(ctx, input.Uuids)

	return err
}
