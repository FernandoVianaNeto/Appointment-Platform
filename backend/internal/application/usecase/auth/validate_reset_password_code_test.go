package auth_usecase_test

import (
	auth_usecase "appointment-platform-backend-backend/internal/application/usecase/auth"
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	mock_repository "appointment-platform-backend-backend/test/mocks/domain_repository"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestValidateResetPasswordCodeUsecase_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("should return true if code is valid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		usecase := auth_usecase.NewValidateResetPasswordCodeUsecase(codeRepo)

		codeRepo.EXPECT().
			IsValidCode(ctx, "user@email.com", 123456).
			Return(true, nil)

		input := dto.ValidateResetPasswordCodeInputDto{
			Email: "user@email.com",
			Code:  123456,
		}

		result, err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, domain_response.ValidateResetPasswordCodeResponse{IsValid: true}, result)
	})

	t.Run("should return false if code is invalid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		usecase := auth_usecase.NewValidateResetPasswordCodeUsecase(codeRepo)

		codeRepo.EXPECT().
			IsValidCode(ctx, "user@email.com", gomock.Any()).
			Return(false, nil)

		input := dto.ValidateResetPasswordCodeInputDto{
			Email: "user@email.com",
			Code:  123456,
		}

		result, err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, domain_response.ValidateResetPasswordCodeResponse{IsValid: false}, result)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		usecase := auth_usecase.NewValidateResetPasswordCodeUsecase(codeRepo)

		codeRepo.EXPECT().
			IsValidCode(ctx, "user@email.com", 123456).
			Return(false, errors.New("db error"))

		input := dto.ValidateResetPasswordCodeInputDto{
			Email: "user@email.com",
			Code:  123456,
		}

		result, err := usecase.Execute(ctx, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "db error")
		assert.Equal(t, domain_response.ValidateResetPasswordCodeResponse{}, result)
	})
}
