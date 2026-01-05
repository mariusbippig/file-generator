package file

import "fmt"

type File interface {
	Write(w []byte) (int, error)
	Close()
	GetFilename() string
}

func NewFile(driver string, filename string, fileType string) (File, error) {
	switch driver {
	case "local":
		return NewLocalFile(filename, fileType)
	default:
		return nil, fmt.Errorf("File driver %s has not been implemented", driver)
	}
}
