package file

import "fmt"

type File interface {
	Write(w []byte) (int, error)
	Close()
	GetFilename() string
}

// NewFile is a factory function which creates a new File instance based on the specified driver.
func NewFile(driver string, filename string, fileType string) (File, error) {
	switch driver {
	case "local":
		return NewLocalFile(filename, fileType)
	default:
		return nil, fmt.Errorf("File driver %s has not been implemented", driver)
	}
}
