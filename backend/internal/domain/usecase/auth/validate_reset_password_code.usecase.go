package domain_auth_usecase

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"context"
)

type ValidateResetPasswordCodeUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ValidateResetPasswordCodeInputDto) (domain_response.ValidateResetPasswordCodeResponse, error)
}
