package main

import (
	"bufio"
	"cmd/pkg/config"
	"cmd/pkg/file"
	"cmd/pkg/logger"
	"fmt"
	"os"
)

func main() {
	config, err := config.NewConfig("app")
	if err != nil {
		panic(err)
	}

	logger, err := logger.NewLogger(config)
	if err != nil {
		panic(err)
	}

	contentGenerator, err := file.NewContentGenerator(config.FileSizeBytes)
	if err != nil {
		logger.Panic(err.Error())
	}

	// TODO replace "file" with driver based file location to have e.g. different storage points
	filename, err := file.EvaluateFilename(config.Filename, config.FileType)
	if err != nil {
		logger.Panic(err.Error())
	}

	// TODO implement creation of multiple files in parallel
	file, err := os.Create(filename)
	if err != nil {
		logger.Panic(err.Error())
	}

	defer file.Close()
	totalFileSize := 0

	bufioWriter := bufio.NewWriter(file)

	for {
		if totalFileSize >= (config.FileSizeBytes) {
			break
		}

		writtenBytes, err := bufioWriter.Write(contentGenerator.GenerateBatch())
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

	logger.Info(fmt.Sprintf("File %s created with size of %d %s.", filename, totalFileSize, unit))
}
