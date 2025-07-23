package patient_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_patient "appointment-platform-backend-backend/internal/domain/usecase/patient"
	"context"
)

type DeletePatientUsecase struct {
	PatientRepository domain_repository.PatientRepositoryInterface
}

func NewDeletePatientUseCase(
	repository domain_repository.PatientRepositoryInterface,
) domain_usecase_patient.DeletePatientUsecaseInterface {
	return &DeletePatientUsecase{
		PatientRepository: repository,
	}
}

func (u *DeletePatientUsecase) Execute(ctx context.Context, input dto.DeletePatientInputDto) error {
	u.PatientRepository.DeleteMany(ctx, input.Uuids)

	return nil
}
