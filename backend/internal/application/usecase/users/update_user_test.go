package user_usecase_test

import (
	"context"
	"errors"
	"testing"

	user_usecase "appointment-platform-backend-backend/internal/application/usecase/users"
	"appointment-platform-backend-backend/internal/domain/dto"
	mock_repo "appointment-platform-backend-backend/test/mocks/domain_repository"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateUserUsecase_Execute(t *testing.T) {
	t.Run("should update user successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repo.NewMockUserRepositoryInterface(ctrl)
		usecase := user_usecase.NewUpdateUserUseCase(mockRepo)

		ctx := context.Background()

		name := "Updated name"

		input := dto.UpdateUserInputDto{
			Uuid: "user-uuid-123",
			Name: &name,
		}

		mockRepo.
			EXPECT().
			UpdateByUuid(ctx, input).
			Return(nil)

		err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repo.NewMockUserRepositoryInterface(ctrl)
		usecase := user_usecase.NewUpdateUserUseCase(mockRepo)

		ctx := context.Background()
		name := "Updated name"

		input := dto.UpdateUserInputDto{
			Uuid: "user-uuid-123",
			Name: &name,
		}

		expectedErr := errors.New("failed to update user")

		mockRepo.
			EXPECT().
			UpdateByUuid(ctx, input).
			Return(expectedErr)

		err := usecase.Execute(ctx, input)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}
