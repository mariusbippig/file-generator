# File Generator

A flexible Go application for generating files with configurable sizes and content. The tool is designed with extensibility in mind, featuring pluggable drivers for file storage, logging, and content generation.

## Features

- 📁 Generate files of any specified size
- 🚀 Parallel file generation using configurable worker pool
- 🛑 Graceful shutdown with signal handling (SIGINT, SIGTERM, SIGHUP)
- ⚙️ Configurable via YAML or environment variables
- 🔌 Pluggable driver architecture for extensibility
- 🪵 Flexible logging system with multiple levels
- 🔒 Thread-safe local file creation with RWMutex
- 🔄 Automatic duplicate filename handling (appends numbers like "file (1).txt", "file (2).txt") for default local file driver
- 🧹 Automatic cleanup of incomplete files on shutdown
- ✅ Comprehensive unit tests for core functionality

## Architecture

The project follows a clean architecture with three main packages:

- **config**: Configuration management using Viper
- **file**: File creation and content generation abstraction
- **logger**: Pluggable logging system

### Driver System

The application uses a driver-based architecture for easy extensibility:

- **File Drivers**: Currently supports `local` file system storage
- **Content Generator Drivers**: Currently includes `mock` generator (fills with 'A' characters)
- **Logger Drivers**: Supports `console` logging

## Installation

```bash
git clone <repository-url>
cd file-generator
go mod download
```

## Configuration

Configure the application using either a YAML file (`app.yaml`) or environment variables.

### YAML Configuration

Create an `app.yaml` file in the `cmd/` directory:

```yaml
file:
  filename: "foobar"
  type: "txt"
  sizeBytes: 5000
  amount: 10
logger:
  driver: "console"
  logLevel: "error"
contentGenerator:
  driver: "mock"
```

### Environment Variables

Alternatively, use environment variables:

```bash
export FILENAME="myfile"
export FILETYPE="txt"
export FILE_SIZE=5000
export FILE_AMOUNT=10
export LOG_DRIVER="console"
export LOG_LEVEL="error"
export CONTENT_GENERATOR_DRIVER="mock"
```

### Configuration Options

| Option | Environment Variable | Default | Description |
|--------|---------------------|---------|-------------|
| `file.driver` | `FILE_DRIVER` | `local` | File storage driver |
| `file.filename` | `FILENAME` | `foobar` | Output filename (without extension) |
| `file.type` | `FILETYPE` | `txt` | File extension |
| `file.sizeBytes` | `FILE_SIZE` | `5000` | Target file size in bytes |
| `file.amount` | `FILE_AMOUNT` | `1` | Number of files to generate |
| `generator.routines` | `GENERATOR_ROUTINES` | `5` | Number of worker goroutines |
| `logger.driver` | `LOG_DRIVER` | `console` | Logger driver to use |
| `logger.logLevel` | `LOG_LEVEL` | `error` | Logging level (debug, info, error) |
| `contentGenerator.driver` | `CONTENT_GENERATOR_DRIVER` | `mock` | Content generator driver |

## Usage

### Basic Usage

```bash
cd cmd/
go run main.go
```

### With Custom Configuration

```bash
export FILENAME="testfile"
export FILE_SIZE=10000000    # 10MB
export FILE_AMOUNT=20        # Generate 20 files
go run cmd/main.go
```

Or with a custom YAML file:

```bash
cd cmd/
# Edit app.yaml with your desired settings
go run main.go
```

## Building

```bash
go build -o file-generator cmd/main.go
./file-generator
```

## Project Structure

```
file-generator/
├── cmd/
│   ├── main.go           # Application entry point
│   └── app.yaml          # Configuration file
├── pkg/
│   ├── config/
│   │   └── config.go     # Configuration management
│   ├── file/
│   │   ├── file.go       # File interface and factory
│   │   ├── localFile.go  # Local file driver implementation
│   │   ├── localFile_test.go  # Unit tests for filename handling
│   │   ├── content.go    # Content generator interface
│   │   └── mock.go       # Mock content generator implementation
│   └── logger/
│       ├── logger.go     # Logger interface and factory
│       ├── logger_test.go     # Unit tests for logger factory
│       ├── console.go    # Console logger implementation
│       ├── console_test.go    # Unit tests for console logger
│       └── noop.go       # No-op logger implementation
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## How It Works

1. **Configuration Loading**: Reads configuration from `app.yaml` or environment variables using Viper
2. **Signal Handling**: Sets up context with signal notification for graceful shutdown (SIGINT, SIGTERM, SIGHUP)
3. **Logger Initialization**: Creates a logger instance based on the configured driver
4. **Job Queue**: Creates a buffered channel with all file generation jobs
5. **Worker Pool**: Spawns configurable number of worker goroutines (`generator.routines`)
6. **Content Generator Setup**: Each worker initializes its own content generator
7. **Thread-Safe File Creation**: LocalFile driver uses RWMutex to prevent concurrent filename conflicts
8. **Buffer-Based Writing**: Writes content in 512-byte chunks until target file size is reached
9. **Context Checks**: Workers check for shutdown signals between files and during writes (every 512 bytes)
10. **Automatic Filename Deduplication**: Appends numbers (e.g., "file (1).txt") when filenames conflict
11. **Graceful Shutdown**: On signal, workers stop immediately and clean up incomplete files
12. **Synchronization**: Main thread waits for all workers to complete using sync.WaitGroup

## Test Coverage

- **Logger Package**: Tests for log level mapping
- **File Package**: Tests for filename evaluation and duplicate handling logic of the local file driver

## Future Enhancements

- [ ] Further test coverage
- [ ] Additional job input drivers
- [ ] Additional content generator drivers (random, pattern-based, etc.)
- [ ] Additional file storage drivers (S3, FTP, etc.)
- [ ] Progress bar for large file generation
- [ ] Configurable buffer size for file writing

## Dependencies

- [spf13/viper](https://github.com/spf13/viper) - Configuration management

## Requirements

- Go 1.25.5 or later

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.