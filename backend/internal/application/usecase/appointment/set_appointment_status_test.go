package appointment_usecase_test

import (
	appointment_usecase "appointment-platform-backend-backend/internal/application/usecase/appointment"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/test/mocks/domain_repository"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSetAppointmentStatusUsecase_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("should update appointment status successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockAppointmentRepo := domain_repository.NewMockAppointmentRepositoryInterface(ctrl)

		input := dto.SetAppointmentStatusInputDto{
			Uuid:   "appointment-uuid",
			Status: "confirmed",
		}

		mockAppointmentRepo.
			EXPECT().
			UpdateStatus(ctx, input.Status, input.Uuid).
			Return(nil)

		usecase := appointment_usecase.NewSetAppointmentStatusUsecase(mockAppointmentRepo)

		err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
	})

	t.Run("should return error if repository fails to update status", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockAppointmentRepo := domain_repository.NewMockAppointmentRepositoryInterface(ctrl)

		input := dto.SetAppointmentStatusInputDto{
			Uuid:   "appointment-uuid",
			Status: "cancelled",
		}

		mockAppointmentRepo.
			EXPECT().
			UpdateStatus(ctx, input.Status, input.Uuid).
			Return(errors.New("update failed"))

		usecase := appointment_usecase.NewSetAppointmentStatusUsecase(mockAppointmentRepo)

		err := usecase.Execute(ctx, input)

		assert.EqualError(t, err, "update failed")
	})
}
