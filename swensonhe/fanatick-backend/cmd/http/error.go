package main

import (
	"net/http"

	"github.com/swensonhe/fanatick-backend/fanatick"
)

// ErrorWriterFunc is a function that writes an error.
type ErrorWriterFunc func(e error)

var statusCodes = map[string]int{
	string(fanatick.ErrorNotFound): http.StatusNotFound,
	string(fanatick.ErrorInternal): http.StatusInternalServerError,
	string(fanatick.ErrorUnauthorized): http.StatusUnauthorized,
}

// NewErrorWriter returns a new error writer.
func NewErrorWriter(w http.ResponseWriter) ErrorWriterFunc {
	return func(e error) {
		NewJSONWriter(w).Write(e, statusCode(e))
	}
}

// Write performs the error writer func
func (f ErrorWriterFunc) Write(e error) {
	f(e)
}

func statusCode(e error) int {
	if code, ok := statusCodes[e.Error()]; ok {
		return code
	}
	return http.StatusInternalServerError
}
