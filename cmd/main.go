package main

import (
	"cmd/pkg/config"
	"cmd/pkg/file"
	"cmd/pkg/logger"
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

// main is the entry point of the application that generates files with specified sizes
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	config, err := config.NewConfig("app")
	if err != nil {
		panic(err)
	}

	logger, err := logger.NewLogger(config)
	if err != nil {
		panic(err)
	}

	createFilesChannel := make(chan int, 1000)
	for i := 0; i < config.File.Amount; i++ {
		createFilesChannel <- i
	}
	close(createFilesChannel)

	var waitGroup sync.WaitGroup

	for range config.Generator.Routines {
		waitGroup.Add(1)
		go func() {
			generateFile(ctx, logger, config, createFilesChannel)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	// FIXME fmt.Printf("%d file(s) created.\n", config.File.Amount)
}

func generateFile(ctx context.Context, logger logger.Logger, config config.Config, createFilesChannel chan int) {
fileLoop:
	for {
		select {
		case <-ctx.Done():
			logger.Info("Shutting down gracefully.")
			return
		case _, ok := <-createFilesChannel:
			if !ok {
				return
			}
		}

		contentGenerator, err := file.NewContentGenerator(config.ContentGenerator.Driver)
		if err != nil {
			logger.Critical(fmt.Sprintf("Initiation of content generator failed: %v", err.Error()))
			return
		}

		file, err := file.NewFile(config.File.Driver, config.File.Filename, config.File.Type)
		if err != nil {
			logger.Critical(fmt.Sprintf("Creation of new file failed: %v", err.Error()))
			continue
		}

		totalFileSize := 0
		bufferSize := 512

		for {
			select {
			case <-ctx.Done():
				logger.Info("Shutting down gracefully.")
				file.Close()
				file.Delete()
				return
			default:
				// Continue processing
			}

			if totalFileSize == config.File.SizeBytes {
				break
			}

			if totalFileSize+bufferSize > config.File.SizeBytes {
				bufferSize = config.File.SizeBytes - totalFileSize
			}

			buffer := make([]byte, bufferSize)

			_, err := contentGenerator.Read(buffer)
			if err != nil {
				logger.Critical(fmt.Sprintf("Reading from content generator failed: %v", err.Error()))
				file.Close()
				file.Delete()
				continue fileLoop
			}

			writtenBytes, err := file.Write(buffer)
			if err != nil {
				logger.Critical(fmt.Sprintf("Writing to file failed: %v", err.Error()))
				file.Close()
				file.Delete()
				continue fileLoop
			}

			totalFileSize += writtenBytes
		}

		file.Close()

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
	}
}
