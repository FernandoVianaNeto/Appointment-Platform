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

func TestEditAppointmentUsecase_Execute(t *testing.T) {
	t.Run("should edit appointment successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockAppointmentRepositoryInterface(ctrl)
		ctx := context.Background()

		startDate := "2025-08-01T09:00:00Z"
		endDate := "2025-08-01T10:00:00Z"
		procedure := "Raio-X"

		input := dto.EditAppointmentInputDto{
			Uuid:      "appointment-uuid",
			StartDate: &startDate,
			EndDate:   &endDate,
			Procedure: &procedure,
		}

		mockRepo.
			EXPECT().
			Edit(ctx, input).
			Return(nil)

		usecase := appointment_usecase.NewEditAppointmentUseCase(mockRepo)

		err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockAppointmentRepositoryInterface(ctrl)
		ctx := context.Background()

		input := dto.EditAppointmentInputDto{
			Uuid: "invalid-uuid",
		}

		mockRepo.
			EXPECT().
			Edit(ctx, input).
			Return(errors.New("failed to update appointment"))

		usecase := appointment_usecase.NewEditAppointmentUseCase(mockRepo)

		err := usecase.Execute(ctx, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "failed to update appointment")
	})
}
