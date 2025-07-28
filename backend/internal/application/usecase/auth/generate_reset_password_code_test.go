package auth_usecase_test

import (
	auth_usecase "appointment-platform-backend-backend/internal/application/usecase/auth"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	"appointment-platform-backend-backend/test/mocks/adapter"
	"appointment-platform-backend-backend/test/mocks/domain_repository"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGenerateResetPasswordCodeUsecase(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := domain_repository.NewMockUserRepositoryInterface(ctrl)
		mockResetRepo := domain_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		mockEmail := adapter.NewMockEmailSenderAdapterInterface(ctrl)

		mockUserRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "user@email.com", "local").
			Return(nil, nil)

		usecase := auth_usecase.NewGenerateResetPasswordCodeUsecase(mockResetRepo, mockUserRepo, mockEmail)

		err := usecase.Execute(ctx, dto.GenerateResetPasswordCodeInputDto{
			Email: "user@email.com",
		})

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("should not generate code if already active", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := domain_repository.NewMockUserRepositoryInterface(ctrl)
		mockResetRepo := domain_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		mockEmail := adapter.NewMockEmailSenderAdapterInterface(ctrl)

		mockUser := &entity.User{
			Uuid:  "uuid-1",
			Email: "user@email.com",
			Name:  "User",
		}

		mockUserRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "user@email.com", "local").
			Return(mockUser, nil)

		mockResetRepo.EXPECT().
			FindActive(ctx, "user@email.com").
			Return(1, nil)

		usecase := auth_usecase.NewGenerateResetPasswordCodeUsecase(mockResetRepo, mockUserRepo, mockEmail)

		err := usecase.Execute(ctx, dto.GenerateResetPasswordCodeInputDto{
			Email: "user@email.com",
		})

		assert.NoError(t, err)
	})

	t.Run("should generate and send reset code successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := domain_repository.NewMockUserRepositoryInterface(ctrl)
		mockResetRepo := domain_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		mockEmail := adapter.NewMockEmailSenderAdapterInterface(ctrl)

		mockUser := &entity.User{
			Uuid:  "uuid-1",
			Email: "user@email.com",
			Name:  "User",
		}

		mockUserRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "user@email.com", "local").
			Return(mockUser, nil)

		mockResetRepo.EXPECT().
			FindActive(ctx, "user@email.com").
			Return(0, nil)

		mockResetRepo.EXPECT().
			Create(ctx, gomock.Any()).
			Return(123456, nil)

		mockEmail.EXPECT().
			SendResetPasswordEmail(ctx, "user@email.com", 123456).
			Return(nil)

		usecase := auth_usecase.NewGenerateResetPasswordCodeUsecase(mockResetRepo, mockUserRepo, mockEmail)

		err := usecase.Execute(ctx, dto.GenerateResetPasswordCodeInputDto{
			Email: "user@email.com",
		})

		assert.NoError(t, err)
	})

	t.Run("should return error if Create fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := domain_repository.NewMockUserRepositoryInterface(ctrl)
		mockResetRepo := domain_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		mockEmail := adapter.NewMockEmailSenderAdapterInterface(ctrl)

		mockUser := &entity.User{
			Uuid:  "uuid-1",
			Email: "user@email.com",
			Name:  "User",
		}

		mockUserRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "user@email.com", "local").
			Return(mockUser, nil)

		mockResetRepo.EXPECT().
			FindActive(ctx, "user@email.com").
			Return(0, nil)

		mockResetRepo.EXPECT().
			Create(ctx, gomock.Any()).
			Return(0, errors.New("create failed"))

		usecase := auth_usecase.NewGenerateResetPasswordCodeUsecase(mockResetRepo, mockUserRepo, mockEmail)

		err := usecase.Execute(ctx, dto.GenerateResetPasswordCodeInputDto{
			Email: "user@email.com",
		})

		assert.Error(t, err)
		assert.Equal(t, "create failed", err.Error())
	})

	t.Run("should return error if email sender fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockUserRepo := domain_repository.NewMockUserRepositoryInterface(ctrl)
		mockResetRepo := domain_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		mockEmail := adapter.NewMockEmailSenderAdapterInterface(ctrl)

		mockUser := &entity.User{
			Uuid:  "uuid-1",
			Email: "user@email.com",
			Name:  "User",
		}

		mockUserRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "user@email.com", "local").
			Return(mockUser, nil)

		mockResetRepo.EXPECT().
			FindActive(ctx, "user@email.com").
			Return(0, nil)

		mockResetRepo.EXPECT().
			Create(ctx, gomock.Any()).
			Return(123456, nil)

		mockEmail.EXPECT().
			SendResetPasswordEmail(ctx, "user@email.com", 123456).
			Return(errors.New("email fail"))

		usecase := auth_usecase.NewGenerateResetPasswordCodeUsecase(mockResetRepo, mockUserRepo, mockEmail)

		err := usecase.Execute(ctx, dto.GenerateResetPasswordCodeInputDto{
			Email: "user@email.com",
		})

		assert.Error(t, err)
		assert.Equal(t, "email fail", err.Error())
	})
}
