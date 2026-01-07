package main

import (
	"cmd/pkg/config"
	"cmd/pkg/file"
	"cmd/pkg/logger"
	"fmt"
	"sync"
)

// main is the entry point of the application that generates files with specified sizes
func main() {
	config, err := config.NewConfig("app")
	if err != nil {
		panic(err)
	}

	logger, err := logger.NewLogger(config)
	if err != nil {
		panic(err)
	}

	var waitGroup sync.WaitGroup

	for i := 0; i < config.File.Amount; i++ {
		waitGroup.Go(func() {
			contentGenerator, err := file.NewContentGenerator(config.ContentGenerator.Driver)
			if err != nil {
				logger.Panic(err.Error())
			}

			file, err := file.NewFile(config.File.Driver, config.File.Filename, config.File.Type)
			if err != nil {
				logger.Panic(err.Error())
			}

			defer file.Close()

			totalFileSize := 0
			defaultBufferSize := 512

			for {
				if totalFileSize == config.File.SizeBytes {
					break
				}

				if totalFileSize+defaultBufferSize > config.File.SizeBytes {
					defaultBufferSize = config.File.SizeBytes - totalFileSize
				}

				buffer := make([]byte, defaultBufferSize)

				_, err := contentGenerator.Read(buffer)
				if err != nil {
					logger.Panic(err.Error())
				}

				writtenBytes, err := file.Write(buffer)
				if err != nil {
					logger.Panic(err.Error())
				}

				totalFileSize += writtenBytes
			}

			unit := "bytes"

			if totalFileSize > 1000000000 {
				totalFileSize = totalFileSize / 1000000000
				unit = "GB"
			} else if totalFileSize > 1000000 {
				totalFileSize = totalFileSize / 1000000
				unit = "MB"
			} else if totalFileSize > 1000 {
				totalFileSize = totalFileSize / 1000
				unit = "KB"
			}

			logger.Debug(fmt.Sprintf("File %s created with size of %d %s.", file.GetFilename(), totalFileSize, unit))
		})
	}

	waitGroup.Wait()

	fmt.Printf("%d file(s) created.\n", config.File.Amount)
}
