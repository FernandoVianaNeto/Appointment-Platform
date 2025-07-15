package exceptions

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToMeta(t *testing.T) {
	tests := []struct {
		name     string
		meta     ErrorMetadataLog
		expected gin.H
	}{
		{
			name: "Should convert ErrorMetadataLog to map correctly",
			meta: ErrorMetadataLog{
				ErrorMetadata: ErrorMetadata{
					"key1": "value1",
					"key2": "value2",
				},
			},
			expected: gin.H{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.meta.ToMeta()
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestToLogError(t *testing.T) {
	tests := []struct {
		name     string
		meta     ErrorMetadataLog
		expected error
	}{
		{
			name: "Should return error from ErrorMetadataLog",
			meta: ErrorMetadataLog{
				error: NewApplicationError(errors.New("error")),
			},
			expected: NewApplicationError(errors.New("error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.meta.ToLogError()
			assert.Equal(t, tt.expected, got)
		})
	}
}
