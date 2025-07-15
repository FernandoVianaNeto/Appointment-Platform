package domain_repository

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"appointment-platform-backend-backend/internal/domain/entity"
	"context"
)

const ResetPasswordCode = "reset_password_code"

type ResetPasswordCodeRepositoryInterface interface {
	Create(ctx context.Context, input entity.ResetPasswordCode) (int, error)
	FindActive(ctx context.Context, email string) (int, error)
	IsValidCode(ctx context.Context, email string, code int) (bool, error)
	ActivateCode(ctx context.Context, email string, code int) error
}
