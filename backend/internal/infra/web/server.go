package web

import (
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	domain_auth_usecase "appointment-platform-backend-backend/internal/domain/usecase/auth"
	domain_usecase_patient "appointment-platform-backend-backend/internal/domain/usecase/patient"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/user"
	"context"

	"github.com/gin-contrib/cors"

	gin "github.com/gin-gonic/gin"
)

type Server struct {
	router                           *gin.Engine
	CreateUserUsecase                domain_usecase.CreateUserUsecaseInterface
	GetUserUsecase                   domain_usecase.GetUserProfileUsecaseInterface
	UpdateUserUsecase                domain_usecase.UpdateUserUsecaseInterface
	AuthUseCase                      domain_auth_usecase.AuthUsecaseInterface
	GoogleAuthUsecase                domain_auth_usecase.GoogleAuthUsecaseInterface
	GenerateResetPasswordCodeUsecase domain_auth_usecase.GenerateResetPasswordCodeUsecaseInterface
	ResetPasswordUsecase             domain_auth_usecase.ResetPasswordUsecaseInterface
	ValidateResetPasswordCodeUsecase domain_auth_usecase.ValidateResetPasswordCodeUsecaseInterface
	ListPatientUsecase               domain_usecase_patient.ListPatientUsecaseInterface
	EditPatientUsecase               domain_usecase_patient.EditPatientUsecaseInterface
	DeletePatientUsecase             domain_usecase_patient.DeletePatientUsecaseInterface
	CreatePatientUsecase             domain_usecase_patient.CreatePatientUsecaseInterface
	CreateAppointmentUsecase         domain_usecase_appointment.CreateAppointmentUsecaseInterface
	EditAppointmentUsecase           domain_usecase_appointment.EditAppointmentUsecaseInterface
	ListAppointmentUsecase           domain_usecase_appointment.ListAppointmentsUsecaseInterface
	DeleteAppointmentUsecase         domain_usecase_appointment.DeleteAppointmentUsecaseInterface
	SetAppointmentStatusUsecase      domain_usecase_appointment.SetAppointmentStatusUsecaseInterface
}

func NewServer(
	ctx context.Context,
	createUserUsecase domain_usecase.CreateUserUsecaseInterface,
	getUserUsecase domain_usecase.GetUserProfileUsecaseInterface,
	updateUserUsecase domain_usecase.UpdateUserUsecaseInterface,
	authUsecase domain_auth_usecase.AuthUsecaseInterface,
	googleAuthUsecase domain_auth_usecase.GoogleAuthUsecaseInterface,
	generateResetPasswordCodeUsecase domain_auth_usecase.GenerateResetPasswordCodeUsecaseInterface,
	resetPasswordUsecase domain_auth_usecase.ResetPasswordUsecaseInterface,
	validateResetPasswordCodeUsecase domain_auth_usecase.ValidateResetPasswordCodeUsecaseInterface,
	listPatientUsecase domain_usecase_patient.ListPatientUsecaseInterface,
	editPatientUsecase domain_usecase_patient.EditPatientUsecaseInterface,
	deletePatientUsecase domain_usecase_patient.DeletePatientUsecaseInterface,
	createPatientUsecase domain_usecase_patient.CreatePatientUsecaseInterface,
	createAppointmentUsecase domain_usecase_appointment.CreateAppointmentUsecaseInterface,
	editAppointmentUsecase domain_usecase_appointment.EditAppointmentUsecaseInterface,
	listAppointmentUsecase domain_usecase_appointment.ListAppointmentsUsecaseInterface,
	deleteAppointmentUsecase domain_usecase_appointment.DeleteAppointmentUsecaseInterface,
	setAppointmentStatusUsecase domain_usecase_appointment.SetAppointmentStatusUsecaseInterface,
) *Server {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	server := &Server{
		CreateUserUsecase:                createUserUsecase,
		GetUserUsecase:                   getUserUsecase,
		UpdateUserUsecase:                updateUserUsecase,
		AuthUseCase:                      authUsecase,
		GoogleAuthUsecase:                googleAuthUsecase,
		GenerateResetPasswordCodeUsecase: generateResetPasswordCodeUsecase,
		ResetPasswordUsecase:             resetPasswordUsecase,
		ValidateResetPasswordCodeUsecase: validateResetPasswordCodeUsecase,
		ListPatientUsecase:               listPatientUsecase,
		EditPatientUsecase:               editPatientUsecase,
		CreatePatientUsecase:             createPatientUsecase,
		DeletePatientUsecase:             deletePatientUsecase,
		CreateAppointmentUsecase:         createAppointmentUsecase,
		EditAppointmentUsecase:           editAppointmentUsecase,
		ListAppointmentUsecase:           listAppointmentUsecase,
		DeleteAppointmentUsecase:         deleteAppointmentUsecase,
		SetAppointmentStatusUsecase:      setAppointmentStatusUsecase,
	}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
