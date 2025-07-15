package app

import (
	"context"

	configs "appointment-platform-backend-backend/cmd/config"
	service "appointment-platform-backend-backend/internal/application/services"
	auth_usecase "appointment-platform-backend-backend/internal/application/usecase/auth"
	user_usecase "appointment-platform-backend-backend/internal/application/usecase/users"
	adapter "appointment-platform-backend-backend/internal/domain/adapters/email_sender"
	"appointment-platform-backend-backend/internal/domain/adapters/messaging"
	storage_adapter "appointment-platform-backend-backend/internal/domain/adapters/storage"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_service "appointment-platform-backend-backend/internal/domain/service"
	domain_auth_usecase "appointment-platform-backend-backend/internal/domain/usecase/auth"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/user"
	"appointment-platform-backend-backend/internal/infra/adapter/minio"
	"appointment-platform-backend-backend/internal/infra/adapter/sendgrid"
	reset_password_code_mongo_repository "appointment-platform-backend-backend/internal/infra/repository/mongo/reset_password_code"
	mongo_repository "appointment-platform-backend-backend/internal/infra/repository/mongo/user"
	"appointment-platform-backend-backend/internal/infra/web"
	mongoPkg "appointment-platform-backend-backend/pkg/mongo"
	natsclient "appointment-platform-backend-backend/pkg/nats"
	"appointment-platform-backend-backend/pkg/storage"

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
}

type Services struct {
	encryptStringService domain_service.EncryptStringServiceInterface
}

type Adapters struct {
	emailSenderAdapter adapter.EmailSenderAdapterInterface
	storageAdapter     storage_adapter.StorageAdapterInterface
}

type Repositories struct {
	UserRepository              domain_repository.UserRepositoryInterface
	ResetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface
}

func NewApplication() *web.Server {
	ctx := context.Background()

	mongoConnectionInput := mongoPkg.MongoInput{
		DSN:      configs.MongoCfg.Dsn,
		Database: configs.MongoCfg.Database,
	}

	db := mongoPkg.NewMongoDatabase(ctx, mongoConnectionInput)

	eventClient := natsclient.New(configs.NatsCfg.Host)
	eventClient.Connect()

	repositories := NewRepositories(ctx, db)

	services := NewServices(ctx)

	adapters := NewAdapters(ctx)

	usecases := NewUseCases(
		ctx,
		repositories.UserRepository,
		repositories.ResetPasswordCodeRepository,
		services,
		adapters,
		eventClient,
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
	)

	return srv
}

func NewRepositories(
	ctx context.Context,
	db *mongo.Database,
) Repositories {
	userRepository := mongo_repository.NewUserRepository(db)
	resetPasswordCodeRepository := reset_password_code_mongo_repository.NewResetPasswordCodeRepository(db)

	return Repositories{
		UserRepository:              userRepository,
		ResetPasswordCodeRepository: resetPasswordCodeRepository,
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
	minioAdapter := NewStorageAdapter(ctx)

	return Adapters{
		emailSenderAdapter: emailSenderAdapter,
		storageAdapter:     minioAdapter,
	}
}

func NewUseCases(
	ctx context.Context,
	userRepository domain_repository.UserRepositoryInterface,
	resetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface,
	services Services,
	adapters Adapters,
	eventClient messaging.Client,
) UseCases {
	userUsecase := user_usecase.NewCreateUserUseCase(userRepository, services.encryptStringService, adapters.storageAdapter)
	getUserUsecase := user_usecase.NewGetUserProfileUseCase(userRepository)
	updateUserUsecase := user_usecase.NewUpdateUserUseCase(userRepository, adapters.storageAdapter)

	//AUTH
	authUsecase := auth_usecase.NewAuthUsecase(userRepository)
	googleAuthUsecase := auth_usecase.NewGoogleAuthUsecase(userRepository)
	generateResetPasswordCodeUsecase := auth_usecase.NewGenerateResetPasswordCodeUsecase(resetPasswordCodeRepository, userRepository, adapters.emailSenderAdapter)
	resetPasswordUsecase := auth_usecase.NewResetPasswordUsecase(userRepository, resetPasswordCodeRepository, services.encryptStringService)
	validateResetPasswordCodeUsecase := auth_usecase.NewValidateResetPasswordCodeUsecase(resetPasswordCodeRepository)

	return UseCases{
		userUseCase:                      userUsecase,
		GetUserUsecase:                   getUserUsecase,
		UpdateUserUsecase:                updateUserUsecase,
		AuthUsecase:                      authUsecase,
		GoogleAuthUsecase:                googleAuthUsecase,
		GenerateResetPasswordCodeUsecase: generateResetPasswordCodeUsecase,
		ResetPasswordUsecase:             resetPasswordUsecase,
		ValidateResetPasswordCodeUsecase: validateResetPasswordCodeUsecase,
	}
}

func NewStorageAdapter(
	ctx context.Context,
) storage_adapter.StorageAdapterInterface {
	client, err := storage.NewMinioClient(
		configs.MinIoCfg.Host,
		configs.MinIoCfg.User,
		configs.MinIoCfg.Password,
	)

	if err != nil {
		panic("Failed to create MinIO client: " + err.Error())
	}

	err = storage.CreateBucketIfNotExists(ctx, client, configs.MinIoCfg.ProfileBucket)

	if err != nil {
		panic("Failed to create profile bucket: " + err.Error())
	}

	err = storage.CreateBucketIfNotExists(ctx, client, configs.MinIoCfg.ProductBucket)

	if err != nil {
		panic("Failed to create product bucket: " + err.Error())
	}

	minioAdapter := minio.NewMinIoAdapter(ctx, client)

	return minioAdapter
}
