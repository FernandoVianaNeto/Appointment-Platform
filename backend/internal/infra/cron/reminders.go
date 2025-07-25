package cron

import (
	adapter "appointment-platform-backend-backend/internal/domain/adapters/email_sender"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	"context"
	"log"

	"github.com/robfig/cron/v3"
)

func StartReminderScheduler(
	ctx context.Context,
	usecase domain_usecase_appointment.GetNextAppointmentsAndSendReminder,
	appointment_repository domain_repository.AppointmentRepositoryInterface,
	emailSenderAdapter adapter.EmailSenderAdapterInterface,
) error {
	c := cron.New()

	c.AddFunc("@every 1m", func() {
		err := checkAppointmentsAndSendReminders()
		if err != nil {
			log.Printf("Erro ao enviar lembretes: %v", err)
		}
	})

	c.Start()
	select {} // Keep it running
}
