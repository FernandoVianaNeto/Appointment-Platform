package patient_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_patient "appointment-platform-backend-backend/internal/domain/usecase/patient"
	"context"
)

type EditPatientUsecase struct {
	PatientRepository domain_repository.PatientRepositoryInterface
}

func NewEditPatientUseCase(
	repository domain_repository.PatientRepositoryInterface,
) domain_usecase_patient.EditPatientUsecaseInterface {
	return &EditPatientUsecase{
		PatientRepository: repository,
	}
}

func (u *EditPatientUsecase) Execute(ctx context.Context, input dto.EditPatientInputDto) error {
	err := u.PatientRepository.Edit(ctx, input)

	return err
}
