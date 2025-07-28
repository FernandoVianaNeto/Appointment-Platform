package patient_usecase_test

import (
	patient_usecase "appointment-platform-backend-backend/internal/application/usecase/patient"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_repository "appointment-platform-backend-backend/test/mocks/domain_repository"
)

func TestListPatientUsecase_Execute(t *testing.T) {
	t.Run("should return list of patients with metadata", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPatientRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)

		usecase := patient_usecase.NewListPatientUseCase(mockPatientRepo)

		ctx := context.Background()

		searchInput := "Maria"

		input := dto.ListPatientInputDto{
			Page:        1,
			SearchInput: &searchInput,
			UserUuid:    "user-uuid",
		}

		address := "Rua A, 123"
		patients := []entity.Patient{
			{
				Uuid:      "patient-uuid-1",
				Name:      "Maria",
				Insurance: "Unimed",
				Phone:     "11999999999",
				Address:   &address,
				Email:     "maria@email.com",
			},
		}

		mockPatientRepo.
			EXPECT().
			List(ctx, input).
			Return(patients, nil)

		mockPatientRepo.
			EXPECT().
			CountDocuments(ctx, input).
			Return(int64(1), nil)

		result, err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
		assert.Len(t, result.Data, 1)
		assert.Equal(t, "Maria", result.Data[0].Name)
		assert.Equal(t, "Rua A, 123", result.Data[0].Address)
		assert.Equal(t, 1, result.Metadata.CurrentPage)
		assert.Equal(t, 1, result.Metadata.Total)
	})
}
