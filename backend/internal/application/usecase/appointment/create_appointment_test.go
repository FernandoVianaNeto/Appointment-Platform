package appointment_usecase_test

import (
	appointment_usecase "appointment-platform-backend-backend/internal/application/usecase/appointment"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	mock_repository "appointment-platform-backend-backend/test/mocks/domain_repository"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateAppointmentUsecase_Execute(t *testing.T) {
	t.Run("should create appointment successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ctx := context.Background()
		mockAppointmentRepo := mock_repository.NewMockAppointmentRepositoryInterface(ctrl)
		mockPatientRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)

		patientUuid := "35ccbd3a-1ff8-45c9-b3ea-11cf90d2146a"

		input := dto.CreateAppointmentInputDto{
			UserUuid:    "user-uuid",
			StartDate:   "2025-07-29T10:00:00Z",
			EndDate:     "2025-07-29T11:00:00Z",
			PatientUuid: patientUuid,
			Insurance:   "Unimed",
			Technician:  "Maria",
			Location:    "Sala 1",
			Procedure:   "Ressonância",
		}

		patient := entity.Patient{
			Uuid: patientUuid,
			Name: "João da Silva",
		}

		mockPatientRepo.
			EXPECT().
			GetByUuid(ctx, patientUuid).
			Return(patient, nil)

		mockAppointmentRepo.
			EXPECT().
			Create(ctx, gomock.Any()).
			Return(nil)

		usecase := appointment_usecase.NewCreateAppointmentUseCase(mockAppointmentRepo, mockPatientRepo)

		err := usecase.Execute(ctx, input)

		assert.NoError(t, err)
	})

	t.Run("should return error when patient repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ctx := context.Background()
		mockAppointmentRepo := mock_repository.NewMockAppointmentRepositoryInterface(ctrl)
		mockPatientRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)

		input := dto.CreateAppointmentInputDto{
			UserUuid:    "user-uuid",
			StartDate:   "2025-07-29T10:00:00Z",
			EndDate:     "2025-07-29T11:00:00Z",
			PatientUuid: "invalid-uuid",
			Insurance:   "SulAmérica",
			Technician:  "José",
			Location:    "Sala 2",
			Procedure:   "Ultrassom",
		}

		mockPatientRepo.
			EXPECT().
			GetByUuid(ctx, input.PatientUuid).
			Return(entity.Patient{}, errors.New("patient not found"))

		usecase := appointment_usecase.NewCreateAppointmentUseCase(mockAppointmentRepo, mockPatientRepo)

		err := usecase.Execute(ctx, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "patient not found")
	})

	t.Run("should return error when appointment repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ctx := context.Background()
		mockAppointmentRepo := mock_repository.NewMockAppointmentRepositoryInterface(ctrl)
		mockPatientRepo := mock_repository.NewMockPatientRepositoryInterface(ctrl)

		patientUuid := "abc-def"

		input := dto.CreateAppointmentInputDto{
			UserUuid:    "user-uuid",
			StartDate:   "2025-07-29T12:00:00Z",
			EndDate:     "2025-07-29T13:00:00Z",
			PatientUuid: patientUuid,
			Insurance:   "Bradesco",
			Technician:  "Ana",
			Location:    "Sala 3",
			Procedure:   "Tomografia",
		}

		patient := entity.Patient{
			Uuid: patientUuid,
			Name: "Paciente XPTO",
		}

		mockPatientRepo.
			EXPECT().
			GetByUuid(ctx, patientUuid).
			Return(patient, nil)

		mockAppointmentRepo.
			EXPECT().
			Create(ctx, gomock.Any()).
			Return(errors.New("db error"))

		usecase := appointment_usecase.NewCreateAppointmentUseCase(mockAppointmentRepo, mockPatientRepo)

		err := usecase.Execute(ctx, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "db error")
	})
}
