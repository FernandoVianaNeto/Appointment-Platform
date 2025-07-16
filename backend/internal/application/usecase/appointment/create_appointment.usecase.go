package appointment_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	"context"

	"github.com/google/uuid"
)

type CreateAppointmentUsecase struct {
	AppointmentRepository domain_repository.AppointmentRepositoryInterface
}

func NewCreateAppointmentUseCase(
	repository domain_repository.AppointmentRepositoryInterface,
) domain_usecase_appointment.CreateAppointmentUsecaseInterface {
	return &CreateAppointmentUsecase{
		AppointmentRepository: repository,
	}
}

func (u *CreateAppointmentUsecase) Execute(ctx context.Context, input dto.CreateAppointmentInputDto) error {
	appointmentUuid := uuid.New().String()

	entity := entity.NewAppointment(
		appointmentUuid,
		input.UserUuid,
		input.StartDate,
		input.EndDate,
		input.PatientUuid,
		input.Insurance,
		input.Technician,
		input.Location,
		input.Procedure,
	)

	err := u.AppointmentRepository.Create(ctx, *entity)

	return err
}
