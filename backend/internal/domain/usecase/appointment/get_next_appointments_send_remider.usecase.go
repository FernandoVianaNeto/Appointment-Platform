package domain_usecase_appointment

import (
	"context"
)

type GetNextAppointmentsAndSendReminder interface {
	Execute(ctx context.Context) error
}
