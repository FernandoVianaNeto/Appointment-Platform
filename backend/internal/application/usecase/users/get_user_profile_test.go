package user_usecase_test

import (
	"context"
	"errors"
	"testing"

	user_usecase "appointment-platform-backend-backend/internal/application/usecase/users"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"

	mock_repo "appointment-platform-backend-backend/test/mocks/domain_repository"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUserProfileUsecase_Execute(t *testing.T) {
	t.Run("should return user profile successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repo.NewMockUserRepositoryInterface(ctrl)

		usecase := user_usecase.NewGetUserProfileUseCase(mockRepo)

		ctx := context.Background()
		input := dto.GetUserInputDto{
			Uuid: "user-uuid-123",
		}

		expectedUser := &entity.User{
			Uuid:  "user-uuid-123",
			Name:  "Test User",
			Email: "test@example.com",
		}

		mockRepo.
			EXPECT().
			GetByUuid(ctx, input.Uuid).
			Return(expectedUser, nil)

		response, err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, expectedUser.Uuid, response.Uuid)
		assert.Equal(t, expectedUser.Name, response.Name)
		assert.Equal(t, expectedUser.Email, response.Email)
	})

	t.Run("should return nil response if user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repo.NewMockUserRepositoryInterface(ctrl)

		usecase := user_usecase.NewGetUserProfileUseCase(mockRepo)

		ctx := context.Background()
		input := dto.GetUserInputDto{
			Uuid: "non-existent-uuid",
		}

		mockRepo.
			EXPECT().
			GetByUuid(ctx, input.Uuid).
			Return(nil, nil)

		response, err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
		assert.Nil(t, response)
	})

	t.Run("should return error if repository returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repo.NewMockUserRepositoryInterface(ctrl)

		usecase := user_usecase.NewGetUserProfileUseCase(mockRepo)

		ctx := context.Background()
		input := dto.GetUserInputDto{
			Uuid: "some-uuid",
		}

		expectedError := errors.New("repository error")

		mockRepo.
			EXPECT().
			GetByUuid(ctx, input.Uuid).
			Return(nil, expectedError)

		response, err := usecase.Execute(ctx, input)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, response)
	})
}
