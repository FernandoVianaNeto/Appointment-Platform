package appointment_usecase

import (
	configs "appointment-platform-backend-backend/cmd/config"
	adapter "appointment-platform-backend-backend/internal/domain/adapters/email_sender"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	"context"
	"fmt"
	"time"
)

type GetNextAppointmentsAndSendReminder struct {
	AppointmentRepository domain_repository.AppointmentRepositoryInterface
	PatientRepository     domain_repository.PatientRepositoryInterface
	EmailSender           adapter.EmailSenderAdapterInterface
}

func NewGetNextAppointmentsAndSendReminder(
	repository domain_repository.AppointmentRepositoryInterface,
	patientRepository domain_repository.PatientRepositoryInterface,
	emailSender adapter.EmailSenderAdapterInterface,
) domain_usecase_appointment.GetNextAppointmentsAndSendReminder {
	return &GetNextAppointmentsAndSendReminder{
		AppointmentRepository: repository,
		PatientRepository:     patientRepository,
		EmailSender:           emailSender,
	}
}

func (u *GetNextAppointmentsAndSendReminder) Execute(ctx context.Context) error {
	var err error

	window := configs.ReminderCronCfg.Window

	appointments, err := u.AppointmentRepository.GetNextAppointments(ctx, time.Duration(window)*time.Hour)

	fmt.Println("CALLED WITH APPOINTMENTS", appointments)

	if err != nil || appointments == nil {
		fmt.Println(appointments, err)

		return err
	}

	for _, appointment := range *appointments {
		patient, err := u.PatientRepository.GetByUuid(ctx, appointment.PatientUuid)

		if err != nil {
			fmt.Println("ERROR ON GET PATIENT BY UUID", appointment, err)
			return err
		}

		err = u.EmailSender.SendAppointmentReminder(ctx, patient.Email, appointment.Procedure, appointment.Technician, fmt.Sprintf("http://localhost:5173/appointment/confirmation?uuid=%s", appointment.Uuid))

		if err != nil {
			fmt.Println("ERROR ON SEND APPOINTMENT REMINDER", appointment, err)
			return err
		}

		err = u.AppointmentRepository.UpdateReminderSent(ctx, appointment.Uuid)

		if err != nil {
			fmt.Println("ERROR ON UPDATE REMINDER SENT", appointment, err)
		}

		fmt.Println("REMINDER SENT AND SUCCESSFULLY UPDATED", appointment)
	}

	return err
}
