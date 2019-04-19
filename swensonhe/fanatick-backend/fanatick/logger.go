package fanatick

// Logger defines the operations performed by a logger.
type Logger interface {
	Error(args ...interface{})
}
