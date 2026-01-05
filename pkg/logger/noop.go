package logger

func NewNoopLogger() Logger {
	return NoopLogger{}
}

type NoopLogger struct{}

func (nl NoopLogger) Panic(message string) {}
func (nl NoopLogger) Error(message string) {}
func (nl NoopLogger) Info(message string)  {}
func (nl NoopLogger) Debug(message string) {}
