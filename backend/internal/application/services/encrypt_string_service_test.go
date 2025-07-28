package service_test

import (
	service "appointment-platform-backend-backend/internal/application/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestEncryptStringService_EncryptString(t *testing.T) {
	t.Run("should encrypt string successfully", func(t *testing.T) {
		svc := service.NewEncryptStringService()
		password := "securePassword123"

		hashed, err := svc.EncryptString(password, bcrypt.DefaultCost)

		assert.NoError(t, err)
		assert.NotEmpty(t, hashed)

		err = bcrypt.CompareHashAndPassword(hashed, []byte(password))
		assert.NoError(t, err)
	})

}
