package http_error

import (
	"net/http"
)

// ErrBadRequest is the default bad request error
var ErrBadRequest = TypeFromStatus(http.StatusBadRequest)
var ErrInternalServerError = TypeFromStatus(http.StatusInternalServerError)
var ErrUnauthorized = TypeFromStatus(http.StatusUnauthorized)
var ErrForbidden = TypeFromStatus(http.StatusForbidden)

type HttpError struct {
	ID    string
	Type  *HttpErrorType
	Cause error
}

func (err *HttpError) Write(w http.ResponseWriter, r *http.Request) {
	err.WriteWith(defaultWriter(w, r))
}

func (err *HttpError) WriteWith(writer ErrorWriter) {
	writer.Write(err)
}

func NewHttpErrorType(code int, description string) *HttpErrorType {
	return &HttpErrorType{
		StatusCode:  code,
		Description: description,
	}
}

func TypeFromStatus(statusCode int) *HttpErrorType {
	text := http.StatusText(statusCode)
	if text == "" {
		text = "Unknown Error"
	}
	return NewHttpErrorType(statusCode, text)
}

func NewError(errType *HttpErrorType, cause error) *HttpError {
	return &HttpError{
		ID:    RandomErrorID(),
		Type:  errType,
		Cause: cause,
	}
}
