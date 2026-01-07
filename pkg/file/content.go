package file

import "fmt"

type ContentGenerator interface {
	Read(p []byte) (int, error)
}

// NewContentGenerator creates a new ContentGenerator instance based on the specified driver.
// It acts as a factory function that returns the appropriate ContentGenerator implementation.
func NewContentGenerator(driver string) (ContentGenerator, error) {
	switch driver {
	case "mock":
		return NewMockContentGenerator(), nil
	default:
		return nil, fmt.Errorf("Content generator driver %s has not been implemented", driver)
	}
}

