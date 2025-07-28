package domain_usecase_patient

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type DeletePatientUsecaseInterface interface {
	Execute(ctx context.Context, input dto.DeletePatientInputDto) error
}
