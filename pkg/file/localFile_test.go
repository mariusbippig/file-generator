package file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEvaluateFilenameNoConflict(t *testing.T) {
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "testfile")
	fileType := "txt"

	result, err := evaluateFilename(filename, fileType)

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	expected := filename + ".txt"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestEvaluateFilenameWithOneConflict(t *testing.T) {
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "testfile")
	fileType := "txt"

	existingFile := filename + ".txt"
	file, err := os.Create(existingFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()

	result, err := evaluateFilename(filename, fileType)

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	expected := filename + " (1).txt"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestEvaluateFilenameWithMultipleConflicts(t *testing.T) {
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "testfile")
	fileType := "txt"

	files := []string{
		filename + ".txt",
		filename + " (1).txt",
	}

	for _, f := range files {
		file, err := os.Create(f)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", f, err)
		}
		file.Close()
	}

	result, err := evaluateFilename(filename, fileType)

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	expected := filename + " (2).txt"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestEvaluateFilenameWithDifferentFileType(t *testing.T) {
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "testfile")

	txtFile := filename + ".txt"
	file, err := os.Create(txtFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()

	result, err := evaluateFilename(filename, "log")

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	expected := filename + ".log"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestEvaluateFilenameWithGaps(t *testing.T) {
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "testfile")
	fileType := "txt"

	// Create base file and (1) - should find (2) as next
	files := []string{
		filename + ".txt",
		filename + " (1).txt",
	}

	for _, f := range files {
		file, err := os.Create(f)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", f, err)
		}
		file.Close()
	}

	result, err := evaluateFilename(filename, fileType)

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	expected := filename + " (2).txt"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestEvaluateFilenameWithSpecialCharacters(t *testing.T) {
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test-file_name")
	fileType := "txt"

	result, err := evaluateFilename(filename, fileType)

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	expected := filename + ".txt"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
