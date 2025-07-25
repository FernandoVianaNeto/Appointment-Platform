package cron

import (
	adapter "appointment-platform-backend-backend/internal/domain/adapters/email_sender"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	"context"
	"fmt"
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

	c.AddFunc("@every 600s", func() {
		fmt.Println("CRON RUN")
		err := usecase.Execute(ctx)
		if err != nil {
			log.Printf("Erro on send reminder: %v", err)
		}
	})

	fmt.Println("CRON SUCCESSFULLY STARTED")
	c.Start()
	select {}
}
