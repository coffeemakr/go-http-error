package http_error

import "log"

type ErrorLogger interface {
	Log(error *HttpError)
}

type ErrorLogFnc func(error *HttpError)

func (fnc ErrorLogFnc) Log(error *HttpError) {
	fnc(error)
}

var defaultLogger = ErrorLogFnc(func(error *HttpError) {
	log.Printf("Error %s (status %d - description %s): %s", error.ID, error.Type.StatusCode, error.Type.Description, error.Cause)
})
