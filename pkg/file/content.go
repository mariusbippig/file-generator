package file

type ContentGenerator interface {
	GenerateBatch() []byte
}

func NewContentGenerator(fileSize int) (ContentGenerator, error) {
	return NewRoughContentGenerator(fileSize), nil
}

type RoughContentGenerator struct {
	BatchSize int
}

func NewRoughContentGenerator(fileSize int) ContentGenerator {
	batchSizeBytes := 8

	// If a file size is wished in different sizes, we apply different precision levels when writing
	// batches to recude hard disk usage.
	if fileSize > 1000000*1000 {
		batchSizeBytes = 1000000 * 30
	} else if fileSize > 1000000 {
		batchSizeBytes = 100000
	} else if fileSize > 100 {
		batchSizeBytes = 50
	}

	return RoughContentGenerator{
		BatchSize: batchSizeBytes,
	}
}

func (rcg RoughContentGenerator) GenerateBatch() []byte {
	batch := []byte("01234567")

	for len(batch) < rcg.BatchSize {
		extra := []byte("01234567")
		batch = append(batch, extra...)
	}

	return batch
}
