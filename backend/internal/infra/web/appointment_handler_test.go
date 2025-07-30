package web_test

import (
	"appointment-platform-backend-backend/internal/domain/dto"
	"appointment-platform-backend-backend/internal/infra/web"
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mocks para os usecases usados no Server
type mockCreateAppointmentUsecase struct{}

func (m *mockCreateAppointmentUsecase) Execute(ctx context.Context, input dto.CreateAppointmentInputDto) error {
	return nil
}

type mockEditAppointmentUsecase struct{}

func (m *mockEditAppointmentUsecase) Execute(ctx context.Context, input dto.EditAppointmentInputDto) error {
	return nil
}

type mockDeleteAppointmentUsecase struct{}

func (m *mockDeleteAppointmentUsecase) Execute(ctx context.Context, input dto.DeleteAppointmentInputDto) error {
	return nil
}

type mockSetAppointmentStatusUsecase struct{}

func (m *mockSetAppointmentStatusUsecase) Execute(ctx context.Context, input dto.SetAppointmentStatusInputDto) error {
	return nil
}

func TestHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := &web.Server{
		CreateAppointmentUsecase:    &mockCreateAppointmentUsecase{},
		EditAppointmentUsecase:      &mockEditAppointmentUsecase{},
		DeleteAppointmentUsecase:    &mockDeleteAppointmentUsecase{},
		SetAppointmentStatusUsecase: &mockSetAppointmentStatusUsecase{},
	}

	t.Run("Test CreateAppointmentHandler", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		writer.WriteField("start_date", "2025-07-28T10:00")
		writer.WriteField("end_date", "2025-07-28T11:00")
		writer.WriteField("patient_uuid", "patient-uuid-123")
		writer.WriteField("insurance", "insurance-x")
		writer.WriteField("technician", "tech-1")
		writer.WriteField("location", "room-101")
		writer.WriteField("procedure", "procedure-x")
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/appointments", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req
		ctx.Set("user_uuid", "user-uuid-123")

		server.CreateAppointmentHandler(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Test EditAppointmentHandler", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		writer.WriteField("uuid", "uuid-1")
		writer.WriteField("status", "done")
		writer.WriteField("start_date", "2025-07-28T10:00")
		writer.WriteField("end_date", "2025-07-28T11:00")
		writer.WriteField("procedure", "procedure-x")
		writer.Close()

		req := httptest.NewRequest(http.MethodPut, "/appointments", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		server.EditAppointmentHandler(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Test DeleteAppointmentHandler", func(t *testing.T) {
		jsonBody := `{"uuids": ["uuid-1", "uuid-2"]}`
		req := httptest.NewRequest(http.MethodDelete, "/appointments", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		server.DeleteAppointmentHandler(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Test SetAppointmentStatus", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/appointments/status?uuid=uuid-1&status=done", nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = req

		server.SetAppointmentStatus(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
