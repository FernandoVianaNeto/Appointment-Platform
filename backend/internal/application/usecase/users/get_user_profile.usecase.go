package user_usecase

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_repository "appointment-platform-backend-backend/internal/domain/repository"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	domain_usecase "appointment-platform-backend-backend/internal/domain/usecase/user"
	"context"
)

type GetUserProfileUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
}

func NewGetUserProfileUseCase(
	repository domain_repository.UserRepositoryInterface,
) domain_usecase.GetUserProfileUsecaseInterface {
	return &GetUserProfileUsecase{
		UserRepository: repository,
	}
}

func (g *GetUserProfileUsecase) Execute(ctx context.Context, input dto.GetUserInputDto) (*domain_response.GetUserProfileResponse, error) {
	repositoryResponse, err := g.UserRepository.GetByUuid(ctx, input.Uuid)

	if repositoryResponse == nil {
		return nil, err
	}

	response := domain_response.GetUserProfileResponse{
		Uuid:  repositoryResponse.Uuid,
		Name:  repositoryResponse.Name,
		Email: repositoryResponse.Email,
	}

	return &response, err
}
