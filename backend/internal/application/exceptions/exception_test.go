package exceptions

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestNewNotFoundRestError(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected RestError
	}{
		{
			name: "Should create a new NotFoundRestError correctly",
			text: "Not found error",
			expected: &restError{
				msg:  "Not found error",
				code: http.StatusNotFound,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNotFoundRestError(tt.text)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestNewUnauthorizedRestError(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected RestError
	}{
		{
			name: "Should create a new UnauthorizedRestError correctly",
			text: "Unauthorized error",
			expected: &restError{
				msg:  "Unauthorized error",
				code: http.StatusUnauthorized,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUnauthorizedRestError(tt.text)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestNewUnprocessableEntityRestError(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected RestError
	}{
		{
			name: "Should create a new UnprocessableEntityRestError correctly",
			text: "Unprocessable entity error",
			expected: &restError{
				msg:  "Unprocessable entity error",
				code: http.StatusUnprocessableEntity,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUnprocessableEntityRestError(tt.text)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestNewInternalServerRestError(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected RestError
	}{
		{
			name: "Should create a new InternalServerRestError correctly",
			text: "Internal server error",
			expected: &restError{
				msg:  "Internal server error",
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInternalServerRestError(tt.text)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestRestErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		expected gin.H
	}{
		{
			name:     "Should return the correct message",
			expected: gin.H{"message": "Not found error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restError := &restError{
				msg: "Not found error",
			}
			got := restError.Message(nil)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestRestErrorError(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "Should return the correct error",
			expected: "Not found error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			restError := &restError{
				msg: "Not found error",
			}
			got := restError.Error()
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestApplicationErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected gin.H
	}{
		{
			name: "Should return the correct message for RestError",
			err:  NewNotFoundRestError("Not found error"),
			expected: gin.H{
				"message": "Not found error",
			},
		},
		{
			name: "Should return the correct message for generic error",
			err:  errors.New("unexpected error"),
			expected: gin.H{
				"message": "unexpected error",
			},
		},
		{
			name: "Should return the correct message for Unauthorized error",
			err:  NewUnauthorizedRestError("Unauthorized error"),
			expected: gin.H{
				"message": "unauthorized",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appError := NewApplicationError(tt.err)
			got := appError.Message(context.Background())
			assert.Equal(t, tt.expected, got)
		})
	}
}
