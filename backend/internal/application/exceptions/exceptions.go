package exceptions

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type restError struct {
	msg  string
	code int
}

func NewRestError(msg string, code int) RestError {
	return &restError{
		msg:  msg,
		code: code,
	}
}

func NewBadRequestRestError(text string) RestError {
	return NewRestError(text, http.StatusBadRequest)
}

func NewUnauthorizedRestError(text string) RestError {
	return NewRestError(text, http.StatusUnauthorized)
}

func NewNotFoundRestError(text string) RestError {
	return NewRestError(text, http.StatusNotFound)
}

func NewUnprocessableEntityError(text string) RestError {
	return NewRestError(text, http.StatusUnprocessableEntity)
}

func NewUnprocessableEntityRestError(text string) RestError {
	return NewRestError(text, http.StatusUnprocessableEntity)
}

func NewInternalServerRestError(text string) RestError {
	return NewRestError(text, http.StatusInternalServerError)
}

func NewGenericError(status int, text string) RestError {
	return NewRestError(text, status)
}

func (e *restError) Code() int {
	return e.code
}

func (e *restError) Message(ctx context.Context) interface{} {
	return gin.H{
		"message": e.msg,
	}
}

func (e *restError) Error() string {
	return e.msg
}

type applicationError struct {
	err error
	m   ErrorMetadataLog
}

func NewApplicationError(err error) ApplicationError {
	return &applicationError{
		err: err,
	}
}

func (e applicationError) Message(ctx context.Context) interface{} {
	var r interface{}

	code := e.Code()

	protocolCode := ctx.Value("codeTracer")
	switch code {
	case http.StatusInternalServerError:
		r = gin.H{
			"message": "unexpected error",
		}
	case http.StatusUnauthorized:
		r = gin.H{
			"message": "unauthorized",
		}
	default:
		switch t := e.err.(type) {
		case RestError:
			r = t.Message(ctx)
		default:
			r = gin.H{
				"message": e.err.Error(),
			}
		}
	}

	if r, ok := r.(gin.H); ok {
		if err, ok := r["error"].(gin.H); ok {
			if protocolCode != nil {
				err["protocol"] = protocolCode.(string)
			}

		}
		if len(e.m.ErrorMetadata) > 0 {
			r["meta"] = e.m.ToMeta()
		}
	}
	return r
}

func (e *applicationError) Code() int {
	switch t := e.err.(type) {
	case RestError:
		return t.Code()
	default:
		return http.StatusInternalServerError
	}
}

func (e *applicationError) Error() string {
	return e.err.Error()
}
