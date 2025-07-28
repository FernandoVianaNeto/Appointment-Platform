package auth_usecase_test

import (
	auth_usecase "appointment-platform-backend-backend/internal/application/usecase/auth"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	mock_repository "appointment-platform-backend-backend/test/mocks/domain_repository"
	mock_service "appointment-platform-backend-backend/test/mocks/domain_service"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestResetPasswordUsecase_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("should reset password successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock_repository.NewMockUserRepositoryInterface(ctrl)
		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		encryptService := mock_service.NewMockEncryptStringServiceInterface(ctrl)

		usecase := auth_usecase.NewResetPasswordUsecase(userRepo, codeRepo, encryptService)

		user := &entity.User{Uuid: "user-uuid", Email: "test@email.com"}

		userRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "test@email.com", "local").
			Return(user, nil)

		codeRepo.EXPECT().
			IsValidCode(ctx, "test@email.com", 123456).
			Return(true, nil)

		encryptService.EXPECT().
			EncryptString("newPass", 10).
			Return([]byte("encrypted"), nil)

		userRepo.EXPECT().
			UpdatePassword(ctx, dto.UserResetPasswordInputDto{
				Uuid:        "user-uuid",
				NewPassword: []byte("encrypted"),
			}).
			Return(nil)

		codeRepo.EXPECT().
			ActivateCode(ctx, "test@email.com", 123456).
			Return(nil)

		err := usecase.Execute(ctx, dto.ResetPasswordInputDto{
			Email:       "test@email.com",
			Code:        123456,
			NewPassword: "newPass",
		})

		assert.NoError(t, err)
	})

	t.Run("should fail if user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock_repository.NewMockUserRepositoryInterface(ctrl)
		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		encryptService := mock_service.NewMockEncryptStringServiceInterface(ctrl)

		usecase := auth_usecase.NewResetPasswordUsecase(userRepo, codeRepo, encryptService)

		userRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "test@email.com", "local").
			Return(nil, nil)

		err := usecase.Execute(ctx, dto.ResetPasswordInputDto{
			Email:       "test@email.com",
			Code:        123456,
			NewPassword: "newPass",
		})

		assert.EqualError(t, err, "user not found")
	})

	t.Run("should fail if repo returns error on GetByEmailAndAuthProvider", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock_repository.NewMockUserRepositoryInterface(ctrl)
		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		encryptService := mock_service.NewMockEncryptStringServiceInterface(ctrl)

		usecase := auth_usecase.NewResetPasswordUsecase(userRepo, codeRepo, encryptService)

		userRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "test@email.com", "local").
			Return(nil, errors.New("db error"))

		err := usecase.Execute(ctx, dto.ResetPasswordInputDto{
			Email:       "test@email.com",
			Code:        123456,
			NewPassword: "newPass",
		})

		assert.EqualError(t, err, "db error")
	})

	t.Run("should fail if code is invalid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock_repository.NewMockUserRepositoryInterface(ctrl)
		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		encryptService := mock_service.NewMockEncryptStringServiceInterface(ctrl)

		usecase := auth_usecase.NewResetPasswordUsecase(userRepo, codeRepo, encryptService)

		user := &entity.User{Uuid: "user-uuid", Email: "test@email.com"}

		userRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "test@email.com", "local").
			Return(user, nil)

		codeRepo.EXPECT().
			IsValidCode(ctx, "test@email.com", 123456).
			Return(false, nil)

		err := usecase.Execute(ctx, dto.ResetPasswordInputDto{
			Email:       "test@email.com",
			Code:        123456,
			NewPassword: "newPass",
		})

		assert.EqualError(t, err, "invalid reset password code")
	})

	t.Run("should fail if encrypt fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock_repository.NewMockUserRepositoryInterface(ctrl)
		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		encryptService := mock_service.NewMockEncryptStringServiceInterface(ctrl)

		usecase := auth_usecase.NewResetPasswordUsecase(userRepo, codeRepo, encryptService)

		user := &entity.User{Uuid: "user-uuid", Email: "test@email.com"}

		userRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "test@email.com", "local").
			Return(user, nil)

		codeRepo.EXPECT().
			IsValidCode(ctx, "test@email.com", 123456).
			Return(true, nil)

		encryptService.EXPECT().
			EncryptString("newPass", 10).
			Return(nil, errors.New("encrypt error"))

		err := usecase.Execute(ctx, dto.ResetPasswordInputDto{
			Email:       "test@email.com",
			Code:        123456,
			NewPassword: "newPass",
		})

		assert.EqualError(t, err, "encrypt error")
	})

	t.Run("should fail if update password fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock_repository.NewMockUserRepositoryInterface(ctrl)
		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		encryptService := mock_service.NewMockEncryptStringServiceInterface(ctrl)

		usecase := auth_usecase.NewResetPasswordUsecase(userRepo, codeRepo, encryptService)

		user := &entity.User{Uuid: "user-uuid", Email: "test@email.com"}

		userRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "test@email.com", "local").
			Return(user, nil)

		codeRepo.EXPECT().
			IsValidCode(ctx, "test@email.com", 123456).
			Return(true, nil)

		encryptService.EXPECT().
			EncryptString("newPass", 10).
			Return([]byte("encrypted"), nil)

		userRepo.EXPECT().
			UpdatePassword(ctx, gomock.Any()).
			Return(errors.New("update error"))

		err := usecase.Execute(ctx, dto.ResetPasswordInputDto{
			Email:       "test@email.com",
			Code:        123456,
			NewPassword: "newPass",
		})

		assert.EqualError(t, err, "update error")
	})

	t.Run("should fail if activate code fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := mock_repository.NewMockUserRepositoryInterface(ctrl)
		codeRepo := mock_repository.NewMockResetPasswordCodeRepositoryInterface(ctrl)
		encryptService := mock_service.NewMockEncryptStringServiceInterface(ctrl)

		usecase := auth_usecase.NewResetPasswordUsecase(userRepo, codeRepo, encryptService)

		user := &entity.User{Uuid: "user-uuid", Email: "test@email.com"}

		userRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "test@email.com", "local").
			Return(user, nil)

		codeRepo.EXPECT().
			IsValidCode(ctx, "test@email.com", 123456).
			Return(true, nil)

		encryptService.EXPECT().
			EncryptString("newPass", 10).
			Return([]byte("encrypted"), nil)

		userRepo.EXPECT().
			UpdatePassword(ctx, gomock.Any()).
			Return(nil)

		codeRepo.EXPECT().
			ActivateCode(ctx, "test@email.com", 123456).
			Return(errors.New("activate error"))

		err := usecase.Execute(ctx, dto.ResetPasswordInputDto{
			Email:       "test@email.com",
			Code:        123456,
			NewPassword: "newPass",
		})

		assert.EqualError(t, err, "activate error")
	})
}
