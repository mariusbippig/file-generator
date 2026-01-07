package file

import "fmt"

type File interface {
	Write(w []byte) (int, error)
	Close()
	GetFilename() string
}

// NewFile creates a new File instance based on the specified driver.
// It acts as a factory function that returns the appropriate File implementation.
func NewFile(driver string, filename string, fileType string) (File, error) {
	switch driver {
	case "local":
		return NewLocalFile(filename, fileType)
	default:
		return nil, fmt.Errorf("File driver %s has not been implemented", driver)
	}
}
