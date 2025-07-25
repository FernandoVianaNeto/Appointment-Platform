package user_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/user"
	"context"
)

type UpdateUserUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
}

func NewUpdateUserUseCase(
	repository domain_repository.UserRepositoryInterface,
) domain_usecase.UpdateUserUsecaseInterface {
	return &UpdateUserUsecase{
		UserRepository: repository,
	}
}

func (u *UpdateUserUsecase) Execute(ctx context.Context, input dto.UpdateUserInputDto) error {
	err := u.UserRepository.UpdateByUuid(ctx, input)

	return err
}
