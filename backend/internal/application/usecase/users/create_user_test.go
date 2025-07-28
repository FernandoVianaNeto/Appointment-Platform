package user_usecase_test

import (
	user_usecase "appointment-platform-backend-backend/internal/application/usecase/users"
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/domain/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_repo "appointment-platform-backend-backend/test/mocks/domain_repository"
	mock_service "appointment-platform-backend-backend/test/mocks/domain_service"
)

func TestCreateUserUsecase_Execute_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockUserRepositoryInterface(ctrl)
	mockEncryptService := mock_service.NewMockEncryptStringServiceInterface(ctrl)

	usecase := user_usecase.NewCreateUserUseCase(mockRepo, mockEncryptService)

	ctx := context.Background()
	password := "123456"
	encryptedPassword := []byte("encrypted123")

	input := dto.CreateUserInputDto{
		Email:    "user@example.com",
		Name:     "User Name",
		Password: &password,
		Origin:   "local",
	}

	mockRepo.
		EXPECT().
		GetByEmailAndAuthProvider(ctx, input.Email, input.Origin).
		Return(nil, nil)

	mockEncryptService.
		EXPECT().
		EncryptString(password, 10).
		Return(encryptedPassword, nil)

	mockRepo.
		EXPECT().
		Create(ctx, gomock.AssignableToTypeOf(entity.User{})).
		Return(nil)

	err := usecase.Execute(ctx, input)
	assert.NoError(t, err)
}
