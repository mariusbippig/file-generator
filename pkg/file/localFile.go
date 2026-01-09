package file

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type LocalFile struct {
	Filename   string
	File       *os.File
	fileWriter *bufio.Writer
}

var localFileList LocalFileList

type LocalFileList struct {
	List    sync.Map
	RWMutex sync.RWMutex
}

// NewLocalFile creates a new local file with the specified filename and type.
func NewLocalFile(filename string, fileType string) (File, error) {
	localFileList.RWMutex.Lock()
	defer localFileList.RWMutex.Unlock()

	filename, err := evaluateFilename(filename, fileType)
	if err != nil {
		return nil, err
	}

	directory := filepath.Dir(filename)
	if len(directory) > 0 {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			return nil, fmt.Errorf("Failed to create missing directory %s: %v", directory, err)
		}
	}

	newFile, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	writer := bufio.NewWriter(newFile)

	return LocalFile{
		Filename:   filename,
		File:       newFile,
		fileWriter: writer,
	}, nil
}

// Write writes the provided byte slice to the local file using a buffered writer.
// It returns the number of bytes written and any error encountered.
func (lf LocalFile) Write(w []byte) (int, error) {
	return lf.fileWriter.Write(w)
}

// Close closes the underlying file handle.
func (lf LocalFile) Close() {
	lf.File.Close()
}

// Delete removes the local file from the filesystem.
func (lf LocalFile) Delete() {
	os.Remove(lf.Filename)
}

// GetFilename returns the full filename including path and extension.
func (lf LocalFile) GetFilename() string {
	return lf.Filename
}

// evaluateFilename checks if a file already exists and generates a unique filename.
// If a conflict exists, it appends a number in parentheses (e.g., "file (1).txt").
func evaluateFilename(filename string, fileType string) (string, error) {
	// check if file with same filename exists already
	_, err := os.Stat(fmt.Sprintf("%s.%s", filename, fileType))
	if os.IsNotExist(err) {
		return fmt.Sprintf("%s.%s", filename, fileType), nil
	}

	if err != nil {
		return "", err
	}

	localFileList.List.Store(fmt.Sprintf("%s.%s", filename, fileType), true)

	newFilename := fmt.Sprintf("%s (1)", filename)

	for number := 1; true; number++ {
		// check if fallback name already exists
		fileInfoFilenameWithSuffix, _ := os.Stat(fmt.Sprintf("%s.%s", newFilename, fileType))
		if fileInfoFilenameWithSuffix == nil {
			filename = newFilename
			break
		}

		// file seems to exist - extract the number from the existing numbered filename
		parts := strings.Split(fileInfoFilenameWithSuffix.Name(), ".")

		// Extract the number from pattern like "filename (N)"
		// Find the last occurrence of " (" to handle filenames that might contain spaces
		baseName := parts[0]
		lastSuffixIndex := strings.LastIndex(baseName, " (")
		if lastSuffixIndex == -1 {
			return "", fmt.Errorf("Failed to parse numbered suffix from existing file '%s.%s'.", filename, fileType)
		}

		// Extract text between " (" and ")"
		numberWithBrackets := baseName[lastSuffixIndex+2 : len(baseName)-1]

		var err error
		number, err = strconv.Atoi(numberWithBrackets)
		if err != nil {
			return "", fmt.Errorf("Failed to parse number from suffix from existing file '%s.%s'.", filename, fileType)
		}

		number++
		newFilename = fmt.Sprintf("%s (%d)", filename, number)
	}

	return fmt.Sprintf("%s.%s", filename, fileType), nil
}
