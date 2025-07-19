package patient_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	domain_usecase_patient "appointment-platform-backend-backend/internal/domain/usecase/patient"
	"context"
)

type ListPatientUsecase struct {
	PatientRepository domain_repository.PatientRepositoryInterface
}

func NewListPatientUseCase(
	patientRepository domain_repository.PatientRepositoryInterface,
) domain_usecase_patient.ListPatientUsecaseInterface {
	return &ListPatientUsecase{
		PatientRepository: patientRepository,
	}
}

func (u *ListPatientUsecase) Execute(ctx context.Context, input dto.ListPatientInputDto) (domain_response.ListPatientsResponse, error) {
	response := []domain_response.PatientData{}
	defaultMetadata := domain_response.GetMetadataParams(input.Page, 0)

	defaultResponse := domain_response.ListPatientsResponse{
		Data:     []domain_response.PatientData{},
		Metadata: defaultMetadata,
	}

	patients, err := u.PatientRepository.List(ctx, input)

	if err != nil || len(patients) == 0 {
		return defaultResponse, err
	}

	allDocuments, err := u.PatientRepository.CountDocuments(ctx, input)

	if err != nil {
		return defaultResponse, err
	}

	for _, patient := range patients {
		response = append(response, domain_response.PatientData{
			Name:      patient.Name,
			Insurance: *patient.Insurance,
			Phone:     patient.Phone,
		})
	}

	metadata := domain_response.GetMetadataParams(input.Page, allDocuments)

	return domain_response.ListPatientsResponse{
		Data:     response,
		Metadata: metadata,
	}, err
}
