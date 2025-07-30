package web_test

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	domain_response "appointment-platform-backend-backend/internal/domain/response"
	"appointment-platform-backend-backend/internal/infra/web"
	"appointment-platform-backend-backend/internal/infra/web/requests"
	"appointment-platform-backend-backend/test/mocks/domain_auth_usecase"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUsecase := domain_auth_usecase.NewMockAuthUsecaseInterface(ctrl)
	server := &web.Server{AuthUseCase: mockAuthUsecase}

	gin.SetMode(gin.TestMode)

	t.Run("invalid JSON returns 400", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("invalid-json"))

		server.AuthHandler(c)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("usecase error returns 401", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		reqBody := requests.AuthRequest{Email: "email@test.com", Password: "pass"}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		mockAuthUsecase.EXPECT().Execute(gomock.Any(), dto.AuthInputDto{
			Email:    "email@test.com",
			Password: "pass",
		}).Return(domain_response.AuthResponse{}, errors.New("auth failed"))

		server.AuthHandler(c)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("success returns 200 and response", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		reqBody := requests.AuthRequest{Email: "email@test.com", Password: "pass"}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		mockAuthUsecase.EXPECT().Execute(gomock.Any(), dto.AuthInputDto{
			Email:    "email@test.com",
			Password: "pass",
		}).Return(domain_response.AuthResponse{
			Token: "jwt-token",
		}, nil)

		server.AuthHandler(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestGenerateResetPasswordCodeHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGenCodeUsecase := domain_auth_usecase.NewMockGenerateResetPasswordCodeUsecaseInterface(ctrl)
	server := &web.Server{GenerateResetPasswordCodeUsecase: mockGenCodeUsecase}

	gin.SetMode(gin.TestMode)

	t.Run("invalid JSON returns 400", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("invalid-json"))
		server.GenerateResetPasswordCodeHandler(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("usecase error returns 422", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		reqBody := requests.GenerateResetPasswordCodeRequest{Email: "email@test.com"}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))

		mockGenCodeUsecase.EXPECT().Execute(gomock.Any(), dto.GenerateResetPasswordCodeInputDto{
			Email: reqBody.Email,
		}).Return(errors.New("failed to generate code"))

		server.GenerateResetPasswordCodeHandler(c)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("success returns 200", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		reqBody := requests.GenerateResetPasswordCodeRequest{Email: "email@test.com"}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))

		mockGenCodeUsecase.EXPECT().Execute(gomock.Any(), dto.GenerateResetPasswordCodeInputDto{
			Email: reqBody.Email,
		}).Return(nil)

		server.GenerateResetPasswordCodeHandler(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestResetPasswordHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockResetPassUsecase := domain_auth_usecase.NewMockResetPasswordUsecaseInterface(ctrl)
	server := &web.Server{ResetPasswordUsecase: mockResetPassUsecase}

	gin.SetMode(gin.TestMode)

	t.Run("invalid JSON returns 400", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("invalid-json"))
		server.ResetPasswordHandler(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("usecase error returns 422", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		reqBody := requests.ResetPasswordRequest{
			Code:        123456,
			NewPassword: "newpass",
			Email:       "email@test.com",
		}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))

		mockResetPassUsecase.EXPECT().Execute(gomock.Any(), dto.ResetPasswordInputDto{
			Code:        reqBody.Code,
			NewPassword: reqBody.NewPassword,
			Email:       reqBody.Email,
		}).Return(errors.New("reset failed"))

		server.ResetPasswordHandler(c)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("success returns 200", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		reqBody := requests.ResetPasswordRequest{
			Code:        123456,
			NewPassword: "newpass",
			Email:       "email@test.com",
		}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))

		mockResetPassUsecase.EXPECT().Execute(gomock.Any(), dto.ResetPasswordInputDto{
			Code:        reqBody.Code,
			NewPassword: reqBody.NewPassword,
			Email:       reqBody.Email,
		}).Return(nil)

		server.ResetPasswordHandler(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestValidateResetPasswordCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockValidateUsecase := domain_auth_usecase.NewMockValidateResetPasswordCodeUsecaseInterface(ctrl)
	server := &web.Server{ValidateResetPasswordCodeUsecase: mockValidateUsecase}

	gin.SetMode(gin.TestMode)

	t.Run("invalid JSON returns 400", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("invalid-json"))
		server.ValidateResetPasswordCode(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("usecase returns not found error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		reqBody := requests.ValidateResetPasswordCodeRequest{
			Code:  123456,
			Email: "email@test.com",
		}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))

		mockValidateUsecase.EXPECT().Execute(gomock.Any(), dto.ValidateResetPasswordCodeInputDto{
			Code:  reqBody.Code,
			Email: reqBody.Email,
		}).Return(nil, errors.New("mongo: not found"))

		server.ValidateResetPasswordCode(c)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("usecase error returns 422", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		reqBody := requests.ValidateResetPasswordCodeRequest{
			Code:  123456,
			Email: "email@test.com",
		}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))

		mockValidateUsecase.EXPECT().Execute(gomock.Any(), dto.ValidateResetPasswordCodeInputDto{
			Code:  reqBody.Code,
			Email: reqBody.Email,
		}).Return(nil, errors.New("some other error"))

		server.ValidateResetPasswordCode(c)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("success returns 200 and response", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)

		reqBody := requests.ValidateResetPasswordCodeRequest{
			Code:  123456,
			Email: "email@test.com",
		}
		jsonBody, _ := json.Marshal(reqBody)
		c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))

		expectedResponse := map[string]interface{}{"valid": true}
		mockValidateUsecase.EXPECT().Execute(gomock.Any(), dto.ValidateResetPasswordCodeInputDto{
			Code:  reqBody.Code,
			Email: reqBody.Email,
		}).Return(expectedResponse, nil)

		server.ValidateResetPasswordCode(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
