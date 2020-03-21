package http_error

import (
	"errors"
	"fmt"
)

type HttpErrorType struct {
	StatusCode  int
	Description string
}

func (t *HttpErrorType) Cause(cause error) *HttpError {
	return NewError(t, cause)
}

func (t *HttpErrorType) Causef(format string, a ...interface{}) *HttpError {
	return NewError(t, fmt.Errorf(format, a...))
}

func (t *HttpErrorType) CauseString(cause string) *HttpError {
	return NewError(t, errors.New(cause))
}

// WithDescription returns a copy of the error type with the description set to the provided parameter
func (t *HttpErrorType) WithDescription(description string) *HttpErrorType {
	return &HttpErrorType{
		StatusCode:  t.StatusCode,
		Description: description,
	}
}
