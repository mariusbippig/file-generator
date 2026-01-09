package logger

import (
	"testing"
)

func TestMapHumanReadableLogLevelDebug(t *testing.T) {
	level, err := mapHumanReadableLogLevel("debug")

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	if level != logLevelDebug {
		t.Errorf("Expected level %d but got %d", logLevelDebug, level)
	}
}

func TestMapHumanReadableLogLevelInfo(t *testing.T) {
	level, err := mapHumanReadableLogLevel("info")

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	if level != logLevelInfo {
		t.Errorf("Expected level %d but got %d", logLevelInfo, level)
	}
}

func TestMapHumanReadableLogLevelError(t *testing.T) {
	level, err := mapHumanReadableLogLevel("error")

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	if level != logLevelError {
		t.Errorf("Expected level %d but got %d", logLevelError, level)
	}
}

func TestMapHumanReadableLogLevelInvalid(t *testing.T) {
	level, err := mapHumanReadableLogLevel("invalid")

	if err == nil {
		t.Errorf("Expected error but got none")
	}

	if level != logLevelError {
		t.Errorf("Expected level %d but got %d", logLevelError, level)
	}
}

func TestMapHumanReadableLogLevelEmpty(t *testing.T) {
	level, err := mapHumanReadableLogLevel("")

	if err == nil {
		t.Errorf("Expected error but got none")
	}

	if level != logLevelError {
		t.Errorf("Expected level %d but got %d", logLevelError, level)
	}
}
