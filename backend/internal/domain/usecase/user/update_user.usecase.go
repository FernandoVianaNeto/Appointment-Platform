package domain_usecase

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"context"
)

type UpdateUserUsecaseInterface interface {
	Execute(ctx context.Context, input dto.UpdateUserInputDto) error
}
