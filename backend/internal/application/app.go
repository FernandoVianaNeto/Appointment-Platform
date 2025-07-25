package app

import (
	"context"

	configs "appointment-platform-backend-backend/cmd/config"
	service "appointment-platform-backend-backend/internal/application/services"
	appointment_usecase "appointment-platform-backend-backend/internal/application/usecase/appointment"
	auth_usecase "appointment-platform-backend-backend/internal/application/usecase/auth"
	patient_usecase "appointment-platform-backend-backend/internal/application/usecase/patient"
	user_usecase "appointment-platform-backend-backend/internal/application/usecase/users"
	adapter "appointment-platform-backend-backend/internal/domain/adapters/email_sender"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_service "appointment-platform-backend-backend/internal/domain/service"
	domain_usecase_appointment "appointment-platform-backend-backend/internal/domain/usecase/appointment"
	domain_auth_usecase "appointment-platform-backend-backend/internal/domain/usecase/auth"
	domain_usecase_patient "appointment-platform-backend-backend/internal/domain/usecase/patient"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/user"
	"appointment-platform-backend-backend/internal/infra/adapter/sendgrid"
	appointment_mongo_repository "appointment-platform-backend-backend/internal/infra/repository/mongo/appointment"
	patient_mongo_repository "appointment-platform-backend-backend/internal/infra/repository/mongo/patient"
	reset_password_code_mongo_repository "appointment-platform-backend-backend/internal/infra/repository/mongo/reset_password_code"
	mongo_repository "appointment-platform-backend-backend/internal/infra/repository/mongo/user"
	"appointment-platform-backend-backend/internal/infra/web"
	mongoPkg "appointment-platform-backend-backend/pkg/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	UseCases UseCases
}
type UseCases struct {
	userUseCase                      domain_usecase.CreateUserUsecaseInterface
	GetUserUsecase                   domain_usecase.GetUserProfileUsecaseInterface
	UpdateUserUsecase                domain_usecase.UpdateUserUsecaseInterface
	AuthUsecase                      domain_auth_usecase.AuthUsecaseInterface
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

type Services struct {
	encryptStringService domain_service.EncryptStringServiceInterface
}

type Adapters struct {
	emailSenderAdapter adapter.EmailSenderAdapterInterface
}

type Repositories struct {
	UserRepository              domain_repository.UserRepositoryInterface
	ResetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface
	AppointmentRepository       domain_repository.AppointmentRepositoryInterface
	PatientRepository           domain_repository.PatientRepositoryInterface
}

func NewApplication() *web.Server {
	ctx := context.Background()

	mongoConnectionInput := mongoPkg.MongoInput{
		DSN:      configs.MongoCfg.Dsn,
		Database: configs.MongoCfg.Database,
	}

	db := mongoPkg.NewMongoDatabase(ctx, mongoConnectionInput)

	repositories := NewRepositories(ctx, db)

	services := NewServices(ctx)

	adapters := NewAdapters(ctx)

	usecases := NewUseCases(
		ctx,
		repositories.UserRepository,
		repositories.ResetPasswordCodeRepository,
		repositories.PatientRepository,
		repositories.AppointmentRepository,
		services,
		adapters,
	)

	srv := web.NewServer(
		ctx,
		usecases.userUseCase,
		usecases.GetUserUsecase,
		usecases.UpdateUserUsecase,
		usecases.AuthUsecase,
		usecases.GoogleAuthUsecase,
		usecases.GenerateResetPasswordCodeUsecase,
		usecases.ResetPasswordUsecase,
		usecases.ValidateResetPasswordCodeUsecase,
		usecases.ListPatientUsecase,
		usecases.EditPatientUsecase,
		usecases.DeletePatientUsecase,
		usecases.CreatePatientUsecase,
		usecases.CreateAppointmentUsecase,
		usecases.EditAppointmentUsecase,
		usecases.ListAppointmentUsecase,
		usecases.DeleteAppointmentUsecase,
		usecases.SetAppointmentStatusUsecase,
	)

	return srv
}

func NewRepositories(
	ctx context.Context,
	db *mongo.Database,
) Repositories {
	userRepository := mongo_repository.NewUserRepository(db)
	resetPasswordCodeRepository := reset_password_code_mongo_repository.NewResetPasswordCodeRepository(db)
	patientRepository := patient_mongo_repository.NewPatientRepository(db)
	appointmentRepository := appointment_mongo_repository.NewAppointmentRepository(db)

	return Repositories{
		UserRepository:              userRepository,
		ResetPasswordCodeRepository: resetPasswordCodeRepository,
		PatientRepository:           patientRepository,
		AppointmentRepository:       appointmentRepository,
	}
}

func NewServices(
	ctx context.Context,
) Services {
	encryptStringService := service.NewEncryptStringService()

	return Services{
		encryptStringService: encryptStringService,
	}
}

func NewAdapters(
	ctx context.Context,
) Adapters {
	emailSenderAdapter := sendgrid.NewEmailSenderAdapter(ctx)

	return Adapters{
		emailSenderAdapter: emailSenderAdapter,
	}
}

func NewUseCases(
	ctx context.Context,
	userRepository domain_repository.UserRepositoryInterface,
	resetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface,
	patientRepository domain_repository.PatientRepositoryInterface,
	appointmentRepository domain_repository.AppointmentRepositoryInterface,
	services Services,
	adapters Adapters,
) UseCases {
	userUsecase := user_usecase.NewCreateUserUseCase(userRepository, services.encryptStringService)
	getUserUsecase := user_usecase.NewGetUserProfileUseCase(userRepository)
	updateUserUsecase := user_usecase.NewUpdateUserUseCase(userRepository)

	//AUTH
	authUsecase := auth_usecase.NewAuthUsecase(userRepository)
	googleAuthUsecase := auth_usecase.NewGoogleAuthUsecase(userRepository)
	generateResetPasswordCodeUsecase := auth_usecase.NewGenerateResetPasswordCodeUsecase(resetPasswordCodeRepository, userRepository, adapters.emailSenderAdapter)
	resetPasswordUsecase := auth_usecase.NewResetPasswordUsecase(userRepository, resetPasswordCodeRepository, services.encryptStringService)
	validateResetPasswordCodeUsecase := auth_usecase.NewValidateResetPasswordCodeUsecase(resetPasswordCodeRepository)

	// PATIENT
	listPatientUsecase := patient_usecase.NewListPatientUseCase(patientRepository)
	editPatientUsecase := patient_usecase.NewEditPatientUseCase(patientRepository)
	deletePatientUsecase := patient_usecase.NewDeletePatientUseCase(patientRepository)
	createPatientUsecase := patient_usecase.NewCreatePatientUseCase(patientRepository)

	//APPOINTMENT
	listAppointmentUsecase := appointment_usecase.NewListAppointmentUseCase(appointmentRepository, patientRepository)
	editAppointmentUsecase := appointment_usecase.NewEditAppointmentUseCase(appointmentRepository)
	deleteAppointmentUsecase := appointment_usecase.NewDeleteAppointmentUseCase(appointmentRepository)
	createAppointmentUsecase := appointment_usecase.NewCreateAppointmentUseCase(appointmentRepository, patientRepository)
	setAppointmentStatusUsecae := appointment_usecase.NewSetAppointmentStatusUsecase(appointmentRepository)

	return UseCases{
		userUseCase:                      userUsecase,
		GetUserUsecase:                   getUserUsecase,
		UpdateUserUsecase:                updateUserUsecase,
		AuthUsecase:                      authUsecase,
		GoogleAuthUsecase:                googleAuthUsecase,
		GenerateResetPasswordCodeUsecase: generateResetPasswordCodeUsecase,
		ResetPasswordUsecase:             resetPasswordUsecase,
		ValidateResetPasswordCodeUsecase: validateResetPasswordCodeUsecase,
		ListPatientUsecase:               listPatientUsecase,
		EditPatientUsecase:               editPatientUsecase,
		DeletePatientUsecase:             deletePatientUsecase,
		CreatePatientUsecase:             createPatientUsecase,
		CreateAppointmentUsecase:         createAppointmentUsecase,
		EditAppointmentUsecase:           editAppointmentUsecase,
		ListAppointmentUsecase:           listAppointmentUsecase,
		DeleteAppointmentUsecase:         deleteAppointmentUsecase,
		SetAppointmentStatusUsecase:      setAppointmentStatusUsecae,
	}
}
