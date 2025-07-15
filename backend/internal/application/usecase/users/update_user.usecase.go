package user_usecase

import (
	storage_adapter "appointment-platform-backend-backend/internal/domain/adapters/storage"
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/user"
	"context"
)

type UpdateUserUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
	StorageAdapter storage_adapter.StorageAdapterInterface
}

func NewUpdateUserUseCase(
	repository domain_repository.UserRepositoryInterface,
	storage storage_adapter.StorageAdapterInterface,
) domain_usecase.UpdateUserUsecaseInterface {
	return &UpdateUserUsecase{
		UserRepository: repository,
		StorageAdapter: storage,
	}
}

func (u *UpdateUserUsecase) Execute(ctx context.Context, input dto.UpdateUserInputDto) error {
	err := u.UserRepository.UpdateByUuid(ctx, input)

	return err
}
