package file

import (
	"bufio"
	"fmt"
	"os"
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
	fileInfoFilename, _ := os.Stat(fmt.Sprintf("%s.%s", filename, fileType))

	if fileInfoFilename != nil {
		localFileList.List.Store(fmt.Sprintf("%s.%s", filename, fileType), true)

		number := 1
		newFilename := fmt.Sprintf("%s (1)", filename)

		for {
			// check if fallback name already exists
			fileInfoFilenameWithSuffix, _ := os.Stat(fmt.Sprintf("%s.%s", newFilename, fileType))
			if fileInfoFilenameWithSuffix == nil {
				filename = newFilename
				break
			}

			// file seems to exist
			parts := strings.Split(fileInfoFilenameWithSuffix.Name(), ".")

			// take filename itself and check if it already contains a number with brackets as suffix
			numberWithBrackets, found := strings.CutPrefix(parts[0], filename+" ")
			if !found {
				return "", fmt.Errorf("Something went wrong with filename stuff")
			}

			// remove brackets from number suffix
			replacer := strings.NewReplacer("(", "", ")", "")
			numberWithBrackets = replacer.Replace(numberWithBrackets)

			var err error
			number, err = strconv.Atoi(numberWithBrackets)
			if err != nil {
				return "", err
			}

			number++
			newFilename = fmt.Sprintf("%s (%d)", filename, number)
		}
	}

	return fmt.Sprintf("%s.%s", filename, fileType), nil
}
