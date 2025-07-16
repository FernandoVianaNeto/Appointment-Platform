package domain_repository

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	"context"
)

const Patient = "patient"

type PatientRepositoryInterface interface {
	Create(ctx context.Context, input entity.Patient) error
	List(ctx context.Context, input dto.ListPatientInputDto) ([]entity.Patient, error)
	Edit(ctx context.Context, input dto.EditPatientInputDto) error
	Delete(ctx context.Context, input dto.DeletePatientInputDto)
}
