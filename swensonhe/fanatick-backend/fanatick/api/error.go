package api

import (
	"strings"

	"github.com/swensonhe/fanatick-backend/fanatick"
)

// Error represents an error.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

// ErrorNotFound returns a not found error.
func ErrorNotFound(messages ...string) Error {
	return Error{Code: fanatick.ErrorNotFound.Error(), Message: strings.Join(messages, " ")}
}

// ErrorInternal returns an internal error.
func ErrorInternal(messages ...string) Error {
	return Error{Code: fanatick.ErrorInternal.Error(), Message: strings.Join(messages, " ")}
}

// Error implements the error interface.
func (e Error) Error() string {
	return e.Code
}
