package http_error

import (
	"encoding/json"
	"log"
	"net/http"
)

type WriterCreator func(w http.ResponseWriter, r *http.Request) ErrorWriter

type loggingErrorWriter struct {
	Delegate ErrorWriter
	Loggers  []ErrorLogger
}

func (w loggingErrorWriter) Write(httpError *HttpError) {
	for _, logger := range w.Loggers {
		logger.Log(httpError)
	}
	w.Delegate.Write(httpError)
}

func withLoggers(writer ErrorWriter, loggers... ErrorLogger) ErrorWriter {
	return &loggingErrorWriter{
		Delegate: writer,
		Loggers: loggers,
	}
}

var defaultWriter WriterCreator = func(w http.ResponseWriter, r *http.Request) ErrorWriter {
	return withLoggers(&JSONWriter{
		ResponseWriter: w,
	}, defaultLogger)
}

type ErrorWriter interface {
	Write(httpError *HttpError)
}

type JSONWriter struct {
	ResponseWriter http.ResponseWriter
}

type httpErrorDto struct {
	Description string `json:"description"`
	Error       bool   `json:"error"`
	ErrorCode   string `json:"error_code,omitempty"`
}

func (writer JSONWriter) Write(httpError *HttpError) {
	writer.ResponseWriter.WriteHeader(httpError.Type.StatusCode)
	writer.ResponseWriter.Header().Set("Content-Type", "application/json")
	httpErrorDto := httpErrorDto{
		Description: httpError.Type.Description,
		Error:       true,
		ErrorCode:   httpError.ID,
	}
	if encodeErr := json.NewEncoder(writer.ResponseWriter).Encode(httpErrorDto); encodeErr != nil {
		// TODO: Handle error
		log.Println("Failed to encode error as JSON", httpError)
	}
}
