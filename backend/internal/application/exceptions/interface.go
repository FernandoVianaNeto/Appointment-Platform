package exceptions

import "context"

type ApplicationError interface {
	RestError
}

type RestError interface {
	Code() int
	Message(ctx context.Context) interface{}
	Error() string
}
