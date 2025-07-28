package appointment_usecase_test

import (
	appointment_usecase "appointment-platform-backend-backend/internal/application/usecase/appointment"
	"appointment-platform-backend-backend/internal/domain/dto"
	mock_repository "appointment-platform-backend-backend/test/mocks/domain_repository"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteAppointmentUsecase_Execute(t *testing.T) {
	t.Run("should delete appointments successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockAppointmentRepositoryInterface(ctrl)
		ctx := context.Background()

		input := dto.DeleteAppointmentInputDto{
			Uuids: []string{"uuid-1", "uuid-2"},
		}

		mockRepo.
			EXPECT().
			DeleteMany(ctx, input.Uuids).
			Return(nil)

		usecase := appointment_usecase.NewDeleteAppointmentUseCase(mockRepo)

		err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockAppointmentRepositoryInterface(ctrl)
		ctx := context.Background()

		input := dto.DeleteAppointmentInputDto{
			Uuids: []string{"uuid-3"},
		}

		mockRepo.
			EXPECT().
			DeleteMany(ctx, input.Uuids).
			Return(errors.New("delete failed"))

		usecase := appointment_usecase.NewDeleteAppointmentUseCase(mockRepo)

		err := usecase.Execute(ctx, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "delete failed")
	})
}
