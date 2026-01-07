package file

type MockContentGenerator struct {
}

// NewMockContentGenerator creates a new MockContentGenerator instance.
// The mock generator fills buffers with the character 'A' (byte 65).
func NewMockContentGenerator() ContentGenerator {
	return &MockContentGenerator{}
}

// Read fills the provided byte slice with mock content (character 'A').
// It returns the number of bytes written and any error encountered.
func (rcg *MockContentGenerator) Read(p []byte) (int, error) {
	readBytes := 0

	for ; readBytes < len(p); readBytes++ {
		p[readBytes] = byte(65)
	}

	return readBytes, nil
}
