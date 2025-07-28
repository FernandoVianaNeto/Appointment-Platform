package domain_usecase_patient

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
)

type ListPatientUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ListPatientInputDto) (domain_response.ListPatientsResponse, error)
}
