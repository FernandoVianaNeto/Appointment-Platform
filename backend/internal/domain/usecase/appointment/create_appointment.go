package domain_usecase_appointment

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type CreateAppointmentUsecaseInterface interface {
	Execute(ctx context.Context, input dto.CreateAppointmentInputDto) error
}
