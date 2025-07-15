package exceptions

import "errors"

var (
	ErrPlanAlreadyExtended = errors.New("Plan already Extended")
	ErrInvalidFreightId    = errors.New("Invalid freight id")
	ErrReasonInvalid       = errors.New("Invalid reason")
)
