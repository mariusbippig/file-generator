package logger

import "fmt"

type ConsoleLogger struct {
	LogLevel int
}

// NewConsoleLogger creates a new ConsoleLogger instance with the specified log level,
// which prints log messages to standard output.
func NewConsoleLogger(logLevel int) Logger {
	return ConsoleLogger{
		LogLevel: logLevel,
	}
}

// Emergency logs a panic-level message and terminates the program.
func (cl ConsoleLogger) Emergency(message string) {
	panic(fmt.Sprintf("[EMERGENCY] %s", message))
}

// Critical logs a critical-level message if the log level permits.
func (cl ConsoleLogger) Critical(message string) {
	if cl.LogLevel >= logLevelEmergency {
		return
	}

	fmt.Printf("[CRITICAL] %s", message)
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
