package domain_usecase_patient

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
)

type ListPatientUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ListPatientInputDto) (domain_response.ListPatientsResponse, error)
}
