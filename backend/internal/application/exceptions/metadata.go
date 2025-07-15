package exceptions

import (
	"github.com/gin-gonic/gin"
)

const validationKey = "validation"

type ErrorMetadataLog struct {
	ErrorMetadata
	error
}

type ErrorMetadata map[string]interface{}

func (m ErrorMetadataLog) ToMeta() interface{} {
	r := make(gin.H)
	for k, v := range m.ErrorMetadata {
		r[k] = v
	}

	return r
}

func (m ErrorMetadataLog) ToLogError() error {
	return m.error
}

type MetadataOption func(err *ApplicationError, value string)
