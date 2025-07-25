package user_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_service "appointment-platform-backend-backend/internal/domain/service"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/user"
	"context"
	"errors"

	"github.com/google/uuid"
)

type CreateUserUsecase struct {
	UserRepository       domain_repository.UserRepositoryInterface
	EncryptStringService domain_service.EncryptStringServiceInterface
}

func NewCreateUserUseCase(
	repository domain_repository.UserRepositoryInterface,
	encryptStringService domain_service.EncryptStringServiceInterface,
) domain_usecase.CreateUserUsecaseInterface {
	return &CreateUserUsecase{
		UserRepository:       repository,
		EncryptStringService: encryptStringService,
	}
}

func (u *CreateUserUsecase) Execute(ctx context.Context, input dto.CreateUserInputDto) error {
	var (
		encryptedPassword []byte
		err               error
	)

	user, err := u.UserRepository.GetByEmailAndAuthProvider(ctx, input.Email, input.Origin)

	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("user already exists")
	}

	if input.Origin == "local" && input.Password != nil {
		encryptedPassword, err = u.EncryptStringService.EncryptString(*input.Password, 10)
		if err != nil {
			return err
		}
	}

	userUuid := uuid.New().String()

	entity := entity.NewUser(
		userUuid,
		input.Email,
		input.Name,
		&encryptedPassword,
		"local",
		nil,
	)

	err = u.UserRepository.Create(ctx, *entity)

	return err
}
