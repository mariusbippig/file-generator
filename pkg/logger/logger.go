package logger

import (
	"cmd/pkg/config"
	"fmt"
)

const logLevelEmergency = 500
const logLevelError = 400
const logLevelInfo = 200
const logLevelDebug = 100

type Logger interface {
	Emergency(message string)
	Critical(message string)
	Error(emessage string)
	Info(message string)
	Debug(message string)
}

// NewLogger is a factory function which creates a new Logger instance based on the configuration.
func NewLogger(config config.Config) (Logger, error) {
	logLevel, err := mapHumanReadableLogLevel(config.Logger.LogLevel)
	if err != nil {
		return NewNoopLogger(), err
	}

	switch config.Logger.Driver {
	case "noop":
		return NewNoopLogger(), nil
	case "console":
		return NewConsoleLogger(logLevel), nil
	}

	return NewNoopLogger(), fmt.Errorf("Invalid logger driver %s", config.Logger.Driver)
}

// mapHumanReadableLogLevel translates a string e.g. "debug" into the related log level number
func mapHumanReadableLogLevel(level string) (int, error) {
	switch level {
	case "error":
		return logLevelError, nil
	case "debug":
		return logLevelDebug, nil
	case "info":
		return logLevelInfo, nil
	}

	return logLevelError, fmt.Errorf("Could not map %s to log level", level)
}
