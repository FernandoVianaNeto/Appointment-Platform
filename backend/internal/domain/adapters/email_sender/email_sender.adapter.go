package adapter

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"context"
)

type EmailSenderAdapterInterface interface {
	SendResetPasswordEmail(ctx context.Context, toEmail string, code int) error
	SendAppointmentReminder(ctx context.Context, toEmail string, procedure string, doctor string, link string) error
}
