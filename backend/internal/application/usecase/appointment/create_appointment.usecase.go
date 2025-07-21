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
	PatientRepository     domain_repository.PatientRepositoryInterface
}

func NewCreateAppointmentUseCase(
	repository domain_repository.AppointmentRepositoryInterface,
	patientRepository domain_repository.PatientRepositoryInterface,
) domain_usecase_appointment.CreateAppointmentUsecaseInterface {
	return &CreateAppointmentUsecase{
		AppointmentRepository: repository,
		PatientRepository:     patientRepository,
	}
}

func (u *CreateAppointmentUsecase) Execute(ctx context.Context, input dto.CreateAppointmentInputDto) error {
	patient, err := u.PatientRepository.GetByUuid(ctx, input.PatientUuid)

	if err != nil {
		return err
	}

	appointmentUuid := uuid.New().String()

	entity := entity.NewAppointment(
		appointmentUuid,
		input.UserUuid,
		input.StartDate,
		input.EndDate,
		input.PatientUuid,
		patient.Name,
		input.Insurance,
		input.Technician,
		input.Location,
		input.Procedure,
	)

	err = u.AppointmentRepository.Create(ctx, *entity)

	return err
}
