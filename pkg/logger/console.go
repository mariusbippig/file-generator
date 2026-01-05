package logger

import "fmt"

type ConsoleLogger struct {
	LogLevel int
}

func NewConsoleLogger(logLevel int) Logger {
	return ConsoleLogger{
		LogLevel: logLevel,
	}
}

func (cl ConsoleLogger) Panic(message string) {
	panic(fmt.Sprintf("[PANIC] %s", message))
}

func (cl ConsoleLogger) Error(message string) {
	if cl.LogLevel >= logLevelError {
		return
	}

	fmt.Printf("[ERROR] %s", message)
}

func (cl ConsoleLogger) Debug(message string) {
	if cl.LogLevel >= logLevelDebug {
		return
	}

	fmt.Printf("[DEBUG] %s", message)
}

func (cl ConsoleLogger) Info(message string) {
	if cl.LogLevel >= logLevelInfo {
		return
	}

	fmt.Printf("[INFO] %s", message)
}
