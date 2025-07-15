package entity

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"time"
)

type ResetPasswordCode struct {
	UserUuid         string    `json:"user_uuid"`
	Code             int       `json:"code"`
	CodeExpiration   time.Time `json:"code_expiration"`
	AlreadyActivated bool      `json:"already_activated"`
	Email            string    `json:"email"`
}

func NewResetPasswordCode(
	userUuid string,
	email string,
) (*ResetPasswordCode, error) {
	code, err := GenerateNumericCode(6)

	if err != nil {
		return nil, err
	}

	codeInt, err := strconv.Atoi(code)

	entity := &ResetPasswordCode{
		UserUuid:         userUuid,
		Code:             codeInt,
		CodeExpiration:   time.Now().Add(15 * time.Minute),
		AlreadyActivated: false,
		Email:            email,
	}

	return entity, err
}

func GenerateNumericCode(length int) (string, error) {
	const digits = "0123456789"
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		bytes[i] = digits[num.Int64()]
	}
	return string(bytes), nil
}
