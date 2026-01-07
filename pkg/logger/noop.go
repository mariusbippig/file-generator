package logger

// NewNoopLogger creates a new NoopLogger instance that discards all log messages.
// Useful for testing or when logging is disabled.
func NewNoopLogger() Logger {
	return NoopLogger{}
}

type NoopLogger struct{}

// Panic does nothing (no-op implementation).
func (nl NoopLogger) Panic(message string) {}

// Error does nothing (no-op implementation).
func (nl NoopLogger) Error(message string) {}

// Info does nothing (no-op implementation).
func (nl NoopLogger) Info(message string)  {}

// Debug does nothing (no-op implementation).
func (nl NoopLogger) Debug(message string) {}
