package auth_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_service "appointment-platform-backend-backend/internal/domain/service"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/auth"
	"context"
	"errors"
)

type ResetPasswordUsecase struct {
	UserRepository              domain_repository.UserRepositoryInterface
	ResetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface
	EncryptionService           domain_service.EncryptStringServiceInterface
}

func NewResetPasswordUsecase(
	repository domain_repository.UserRepositoryInterface,
	resetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface,
	encryptStringService domain_service.EncryptStringServiceInterface,
) domain_usecase.ResetPasswordUsecaseInterface {
	return &ResetPasswordUsecase{
		UserRepository:              repository,
		ResetPasswordCodeRepository: resetPasswordCodeRepository,
		EncryptionService:           encryptStringService,
	}
}

func (a *ResetPasswordUsecase) Execute(ctx context.Context, input dto.ResetPasswordInputDto) error {
	user, err := a.UserRepository.GetByEmailAndAuthProvider(ctx, input.Email, "local")

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	isValidCode, _ := a.ResetPasswordCodeRepository.IsValidCode(ctx, input.Email, input.Code)

	if !isValidCode {
		return errors.New("invalid reset password code")
	}

	encryptedPassword, err := a.EncryptionService.EncryptString(input.NewPassword, 10)

	if err != nil {
		return err
	}

	err = a.UserRepository.UpdatePassword(ctx, dto.UserResetPasswordInputDto{
		Uuid:        user.Uuid,
		NewPassword: encryptedPassword,
	})

	if err != nil {
		return err
	}

	err = a.ResetPasswordCodeRepository.ActivateCode(ctx, input.Email, input.Code)

	return err
}
