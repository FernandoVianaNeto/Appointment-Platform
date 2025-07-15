package auth_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/auth"
	"context"
)

type ValidateResetPasswordCodeUsecase struct {
	ResetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface
}

func NewValidateResetPasswordCodeUsecase(
	resetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface,
) domain_usecase.ValidateResetPasswordCodeUsecaseInterface {
	return &ValidateResetPasswordCodeUsecase{
		ResetPasswordCodeRepository: resetPasswordCodeRepository,
	}
}

func (a *ValidateResetPasswordCodeUsecase) Execute(ctx context.Context, input dto.ValidateResetPasswordCodeInputDto) (domain_response.ValidateResetPasswordCodeResponse, error) {
	isValidCode, err := a.ResetPasswordCodeRepository.IsValidCode(ctx, input.Email, input.Code)

	if err != nil {
		return domain_response.ValidateResetPasswordCodeResponse{}, err
	}

	return domain_response.ValidateResetPasswordCodeResponse{
		IsValid: isValidCode,
	}, nil
}
