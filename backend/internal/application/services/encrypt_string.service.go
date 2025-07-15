package service

import (
	domain_service "appointment-platform-backend-backend/internal/domain/service"

	"golang.org/x/crypto/bcrypt"
)

type EncryptStringService struct{}

func NewEncryptStringService() domain_service.EncryptStringServiceInterface {
	return &EncryptStringService{}
}

func (p *EncryptStringService) EncryptString(str string, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(str), cost)
}
