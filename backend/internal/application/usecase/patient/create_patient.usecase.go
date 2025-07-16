package patient_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_patient "appointment-platform-backend-backend/internal/domain/usecase/patient"
	"context"

	"github.com/google/uuid"
)

type CreatePatientUsecase struct {
	PatientRepository domain_repository.PatientRepositoryInterface
}

func NewCreatePatientUseCase(
	repository domain_repository.PatientRepositoryInterface,
) domain_usecase_patient.CreatePatientUsecaseInterface {
	return &CreatePatientUsecase{
		PatientRepository: repository,
	}
}

func (u *CreatePatientUsecase) Execute(ctx context.Context, input dto.CreatePatientInputDto) error {
	patientUuid := uuid.New().String()

	entity := entity.NewPatient(
		patientUuid,
		input.Name,
		input.Phone,
		input.Insurance,
		input.Address,
		input.Email,
	)

	err := u.PatientRepository.Create(ctx, *entity)

	return err
}
