package logger

import "fmt"

type ConsoleLogger struct {
	LogLevel int
}

// NewConsoleLogger creates a new ConsoleLogger instance with the specified log level.
// Log messages are printed to standard output.
func NewConsoleLogger(logLevel int) Logger {
	return ConsoleLogger{
		LogLevel: logLevel,
	}
}

// Panic logs a panic-level message and terminates the program.
func (cl ConsoleLogger) Panic(message string) {
	panic(fmt.Sprintf("[PANIC] %s", message))
}

// Error logs an error-level message if the log level permits.
func (cl ConsoleLogger) Error(message string) {
	if cl.LogLevel >= logLevelError {
		return
	}

	fmt.Printf("[ERROR] %s", message)
}

// Debug logs a debug-level message if the log level permits.
func (cl ConsoleLogger) Debug(message string) {
	if cl.LogLevel >= logLevelDebug {
		return
	}

	fmt.Printf("[DEBUG] %s", message)
}

// Info logs an info-level message if the log level permits.
func (cl ConsoleLogger) Info(message string) {
	if cl.LogLevel >= logLevelInfo {
		return
	}

	fmt.Printf("[INFO] %s", message)
}
