package sendgrid

import (
	configs "appointment-platform-backend-backend/cmd/config"
	adapter "appointment-platform-backend-backend/internal/domain/adapters/email_sender"
	"context"
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailSenderAdapter struct {
}

func NewEmailSenderAdapter(ctx context.Context) adapter.EmailSenderAdapterInterface {
	return &EmailSenderAdapter{}
}

func (f *EmailSenderAdapter) SendResetPasswordEmail(ctx context.Context, toEmail string, code int) error {
	from := mail.NewEmail("Appointment Plataform", "forfit.application@gmail.com")
	subject := "Your reset password code is here"
	to := mail.NewEmail("User", "fernando.viana.nt@gmail.com")
	plainTextContent := fmt.Sprintf("Use the following code to reset your password: %d", code)
	htmlContent := fmt.Sprintf("<strong>Use the following code to reset your password: %d</strong>", code)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(configs.SendGridCfg.ApiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return nil
}

func (f *EmailSenderAdapter) SendAppointmentReminder(ctx context.Context, toEmail string, procedure string, doctor string, link string) error {
	from := mail.NewEmail("Appointment Plataform", "forfit.application@gmail.com")
	subject := "Confirm appointment"
	to := mail.NewEmail("User", toEmail)
	plainTextContent := fmt.Sprintf("Click in the link to confirm, reschedule or cancel your appointment: %s", link)
	htmlContent := fmt.Sprintf("<strong>Click in the link to confirm, reschedule or cancel your appointment: %s</strong>", link)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(configs.SendGridCfg.ApiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return nil
}
