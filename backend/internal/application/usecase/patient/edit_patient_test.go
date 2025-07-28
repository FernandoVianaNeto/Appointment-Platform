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

func TestEditPatientUsecase_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("should edit patient successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)
		usecase := patient_usecase.NewEditPatientUseCase(mockRepo)

		name := "Novo Nome"
		phone := "11999999999"
		address := "Rua Nova, 123"
		email := "paciente@email.com"

		input := dto.EditPatientInputDto{
			Uuid:    "uuid-paciente",
			Name:    &name,
			Phone:   &phone,
			Address: &address,
			Email:   &email,
		}

		mockRepo.EXPECT().
			Edit(ctx, input).
			Return(nil)

		err := usecase.Execute(ctx, input)
		assert.NoError(t, err)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)
		usecase := patient_usecase.NewEditPatientUseCase(mockRepo)

		input := dto.EditPatientInputDto{
			Uuid: "uuid-paciente",
		}

		mockRepo.EXPECT().
			Edit(ctx, input).
			Return(errors.New("edit error"))

		err := usecase.Execute(ctx, input)
		assert.Error(t, err)
		assert.EqualError(t, err, "edit error")
	})
}
