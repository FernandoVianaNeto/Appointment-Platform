package auth_usecase_test

import (
	auth_usecase "appointment-platform-backend-backend/internal/application/usecase/auth"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	"appointment-platform-backend-backend/test/mocks/domain_repository"

	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthUsecase_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if user is not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := domain_repository.NewMockUserRepositoryInterface(ctrl)

		mockRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "notfound@email.com", "local").
			Return(nil, nil)

		usecase := auth_usecase.NewAuthUsecase(mockRepo)

		input := dto.AuthInputDto{
			Email:    "notfound@email.com",
			Password: "any",
		}

		_, err := usecase.Execute(ctx, input)
		assert.EqualError(t, err, "user not found")
	})

	t.Run("should return error if repository returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := domain_repository.NewMockUserRepositoryInterface(ctrl)

		mockRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, gomock.Any(), gomock.Any()).
			Return(nil, errors.New("user not found"))

		usecase := auth_usecase.NewAuthUsecase(mockRepo)

		_, err := usecase.Execute(ctx, dto.AuthInputDto{
			Email:    "fail@email.com",
			Password: "123",
		})

		assert.EqualError(t, err, "user not found")
	})

	t.Run("should return error if password is incorrect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := domain_repository.NewMockUserRepositoryInterface(ctrl)

		hashed, _ := bcrypt.GenerateFromPassword([]byte("correct"), bcrypt.DefaultCost)
		mockUser := &entity.User{
			Uuid:     "uuid",
			Email:    "email",
			Password: bytePtrSlice(hashed),
		}

		mockRepo.EXPECT().
			GetByEmailAndAuthProvider(ctx, "email", "local").
			Return(mockUser, nil)

		usecase := auth_usecase.NewAuthUsecase(mockRepo)

		_, err := usecase.Execute(ctx, dto.AuthInputDto{
			Email:    "email",
			Password: "wrong",
		})

		assert.Error(t, err)
	})
}

func bytePtrSlice(b []byte) *[]byte {
	return &b
}
