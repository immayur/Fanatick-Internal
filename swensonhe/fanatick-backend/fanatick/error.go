package fanatick

// Error is a Fanatick error.
type Error string

// Errors
const (
	ErrorNotFound = Error(`not_found`)
	ErrorInternal = Error(`internal`)
	ErrorUnauthorized = Error(`unauthorized`)
)

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}
