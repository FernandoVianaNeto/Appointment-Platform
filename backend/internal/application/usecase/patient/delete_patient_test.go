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

func TestDeletePatientUsecase_Execute(t *testing.T) {
	ctx := context.Background()

	t.Run("should delete patients successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)
		usecase := patient_usecase.NewDeletePatientUseCase(mockRepo)

		input := dto.DeletePatientInputDto{
			Uuids: []string{"uuid1", "uuid2"},
		}

		mockRepo.EXPECT().
			DeleteMany(ctx, input.Uuids).
			Return(nil)

		err := usecase.Execute(ctx, input)
		assert.NoError(t, err)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)
		usecase := patient_usecase.NewDeletePatientUseCase(mockRepo)

		input := dto.DeletePatientInputDto{
			Uuids: []string{"uuid1", "uuid2"},
		}

		mockRepo.EXPECT().
			DeleteMany(ctx, input.Uuids).
			Return(errors.New("delete error"))

		err := usecase.Execute(ctx, input)
		assert.Error(t, err)
		assert.EqualError(t, err, "delete error")
	})
}
