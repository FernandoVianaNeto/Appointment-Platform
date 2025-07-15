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
	from := mail.NewEmail("ForFit App", "forfit.application@gmail.com")
	subject := "Aqui está seu código de redefinição de senha"
	to := mail.NewEmail("User", "fernando.viana.nt@gmail.com")
	plainTextContent := fmt.Sprintf("Utilize o seguinte código para realizar a redefinição da sua senha: %d", code)
	htmlContent := fmt.Sprintf("<strong>Utilize o seguinte código para realizar a redefinição da sua senha: %d</strong>", code)
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
