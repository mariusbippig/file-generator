# File Generator

A flexible Go application for generating files with configurable sizes and content. The tool is designed with extensibility in mind, featuring pluggable drivers for file storage, logging, and content generation.

## Features

- 📁 Generate files of any specified size
- ⚙️ Configurable via YAML or environment variables
- 🔌 Pluggable driver architecture for extensibility
- 📊 Human-readable file size output (bytes, KB, MB, GB)
- 🪵 Flexible logging system

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
  sizeBytes: 2000
  amount: 1
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
export FILE_AMOUNT=1
export LOG_DRIVER="console"
export LOG_LEVEL="debug"
export CONTENT_GENERATOR_DRIVER="mock"
```

### Configuration Options

| Option | Environment Variable | Default | Description |
|--------|---------------------|---------|-------------|
| `file.driver` | `FILE_DRIVER` | `local` | File storage driver |
| `file.filename` | `FILENAME` | `foobar` | Output filename (without extension) |
| `file.type` | `FILETYPE` | `txt` | File extension |
| `file.sizeBytes` | `FILE_SIZE` | `5000` | Target file size in bytes |
| `file.amount` | `FILE_AMOUNT` | `1` | Number of files to generate (not yet implemented) |
| `logger.driver` | `LOG_DRIVER` | `console` | Logger driver to use |
| `logger.logLevel` | `LOG_LEVEL` | `debug` | Logging level |
| `contentGenerator.driver` | `CONTENT_GENERATOR_DRIVER` | `mock` | Content generator driver |

## Usage

### Basic Usage

```bash
cd cmd/
go run main.go
```

### With Custom Configuration

```bash
export FILENAME="testfile.txt"
export FILE_SIZE=10000000  # 10MB
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
│   │   ├── content.go    # Content generator interface
│   │   └── mock.go       # Mock content generator implementation
│   └── logger/
│       ├── logger.go     # Logger interface and factory
│       ├── console.go    # Console logger implementation
│       └── noop.go       # No-op logger implementation
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## How It Works

1. **Configuration Loading**: Reads configuration from `app.yaml` or environment variables using Viper
2. **Logger Initialization**: Creates a logger instance based on the configured driver
3. **Content Generator Setup**: Initializes the content generator (currently supports mock data)
4. **File Creation**: Creates the output file using the specified driver
5. **Buffer-Based Writing**: Writes content in 512-byte chunks until the target file size is reached
6. **Progress Reporting**: Logs the final file size in human-readable format

## Future Enhancements

- [ ] Parallel creation of multiple files (see TODO in main.go)
- [ ] Additional content generator drivers (random, pattern-based, etc.)
- [ ] Additional file storage drivers (S3, FTP, etc.)
- [ ] Progress bar for large file generation
- [ ] Streaming support for very large files

## Dependencies

- [spf13/viper](https://github.com/spf13/viper) - Configuration management

## Requirements

- Go 1.25.5 or later

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
