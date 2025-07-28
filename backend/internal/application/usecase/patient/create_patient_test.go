package patient_usecase_test

import (
	patient_usecase "appointment-platform-backend-backend/internal/application/usecase/patient"
	"appointment-platform-backend-backend/internal/domain/dto"
	mock_repository "appointment-platform-backend-backend/test/mocks/domain_repository"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreatePatientUsecase_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("should create patient successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPatientRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)
		usecase := patient_usecase.NewCreatePatientUseCase(mockPatientRepo)

		address := "456 Elm St"

		input := dto.CreatePatientInputDto{
			Uuid:      "723f3d89-ceca-49e9-ae44-954ba789f8a2",
			UserUuid:  "user-uuid",
			Name:      "Jane Doe",
			Phone:     "987654321",
			Insurance: "Bradesco",
			Address:   &address,
			Email:     "jane@example.com",
		}

		mockPatientRepo.EXPECT().
			Create(ctx, gomock.Any()).
			Return(nil)

		err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPatientRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)
		usecase := patient_usecase.NewCreatePatientUseCase(mockPatientRepo)

		address := "456 Elm St"

		input := dto.CreatePatientInputDto{
			UserUuid:  "user-uuid",
			Name:      "Jane Doe",
			Phone:     "987654321",
			Insurance: "Bradesco",
			Address:   &address,
			Email:     "jane@example.com",
		}

		mockPatientRepo.EXPECT().
			Create(ctx, gomock.Any()).
			Return(errors.New("db error"))

		err := usecase.Execute(ctx, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "db error")
	})
}
