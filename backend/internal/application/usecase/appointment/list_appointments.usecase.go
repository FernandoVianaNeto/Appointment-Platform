package appointment_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	"context"
)

type ListAppointmentUsecase struct {
	AppointmentRepository domain_repository.AppointmentRepositoryInterface
	PatientRepository     domain_repository.PatientRepositoryInterface
}

func NewListAppointmentUseCase(
	repository domain_repository.AppointmentRepositoryInterface,
	patientRepository domain_repository.PatientRepositoryInterface,
) domain_usecase_appointment.ListAppointmentsUsecaseInterface {
	return &ListAppointmentUsecase{
		AppointmentRepository: repository,
		PatientRepository:     patientRepository,
	}
}

func (u *ListAppointmentUsecase) Execute(ctx context.Context, input dto.ListAppointmentInputDto) (domain_response.ListAppointmentsResponse, error) {
	response := []domain_response.AppointmentData{}
	defaultMetadata := domain_response.GetMetadataParams(input.Page, 0)

	defaultResponse := domain_response.ListAppointmentsResponse{
		Data:     []domain_response.AppointmentData{},
		Metadata: defaultMetadata,
	}

	appointments, err := u.AppointmentRepository.List(ctx, input)

	if err != nil {
		return defaultResponse, err
	}

	if len(appointments) == 0 {
		return defaultResponse, nil
	}

	for i, appointment := range appointments {
		patient, err := u.PatientRepository.GetByUuid(ctx, appointment.PatientUuid)

		if err == nil {
			response[i] = domain_response.AppointmentData{
				StartDate: appointment.StartDate,
				EndDate:   appointment.EndDate,
				Patient: domain_response.Patient{
					Name:      patient.Name,
					Insurance: *patient.Insurance,
					Phone:     patient.Phone,
				},
				Procedure: appointment.Procedure,
				Location:  appointment.Location,
				Status:    appointment.Status,
			}
		} else {
			response[i] = domain_response.AppointmentData{}
		}
	}

	allDocuments, err := u.AppointmentRepository.CountDocuments(ctx, input)

	if err != nil {
		return defaultResponse, err
	}

	metada := domain_response.GetMetadataParams(input.Page, allDocuments)

	return domain_response.ListAppointmentsResponse{
		Data:     response,
		Metadata: metada,
	}, err
}
