package domain_usecase_appointment

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"context"
)

type GetNextAppointmentsAndSendReminder interface {
	Execute(ctx context.Context) error
}
