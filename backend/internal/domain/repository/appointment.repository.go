package domain_repository

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	"context"
)

const Appointment = "appointment"

type AppointmentRepositoryInterface interface {
	Create(ctx context.Context, input entity.Appointment) error
	List(ctx context.Context, input dto.ListAppointmentInputDto) ([]entity.Appointment, error)
	Edit(ctx context.Context, input dto.EditAppointmentInputDto) error
	DeleteMany(ctx context.Context, appointmentIds []string) error
	CountDocuments(ctx context.Context, input dto.ListAppointmentInputDto) (int64, error)
}
