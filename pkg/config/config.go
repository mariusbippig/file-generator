package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	File             File
	Logger           Logger
	ContentGenerator ContentGenerator
}

type Logger struct {
	Driver   string
	LogLevel string
}

type ContentGenerator struct {
	Driver string
}

type File struct {
	Driver    string
	Filename  string
	Type      string
	SizeBytes int
	Amount    int
}

// NewConfig initializes a new config by the configured config source.
func NewConfig(path string) (Config, error) {
	// TODO add new config origins as drivers here.
	driver := "env"

	switch driver {
	default:
		return NewConfigFromEnv(path)
	}
}

// NewConfigFromEnv loads configuration from a YAML file and environment variables.
// It uses Viper to read the config file and bind environment variables, returning a Config struct.
func NewConfigFromEnv(path string) (Config, error) {
	viper.SetConfigName(path)
	viper.AddConfigPath(".")

	var fileLookupError viper.ConfigFileNotFoundError
	err := viper.ReadInConfig()

	if !errors.As(err, &fileLookupError) && err != nil {
		return Config{}, err
	}

	viper.SetDefault("file.driver", "local")
	viper.SetDefault("file.filename", "foobar")
	viper.SetDefault("file.type", "txt")
	viper.SetDefault("file.sizeBytes", 5000)
	viper.SetDefault("file.amount", 1)
	viper.SetDefault("logger.driver", "console")
	viper.SetDefault("logger.logLevel", "debug")
	viper.SetDefault("contentGenerator.driver", "mock")

	viper.BindEnv("file.driver", "FILE_DRIVER")
	viper.BindEnv("file.filename", "FILENAME")
	viper.BindEnv("file.type", "FILETYPE")
	viper.BindEnv("file.sizeBytes", "FILE_SIZE")
	viper.BindEnv("file.amount", "FILE_AMOUNT")
	viper.BindEnv("logger.driver", "LOG_DRIVER")
	viper.BindEnv("logger.logLevel", "LOG_LEVEL")
	viper.BindEnv("contentGenerator.driver", "CONTENT_GENERATOR_DRIVER")

	config := Config{
		File: File{
			Driver:    viper.GetString("file.driver"),
			Filename:  viper.GetString("file.filename"),
			Type:      viper.GetString("file.type"),
			SizeBytes: viper.GetInt("file.sizeBytes"),
			Amount:    viper.GetInt("file.amount"),
		},
		Logger: Logger{
			Driver:   viper.GetString("logger.driver"),
			LogLevel: viper.GetString("logger.logLevel"),
		},
		ContentGenerator: ContentGenerator{
			Driver: viper.GetString("contentGenerator.driver"),
		},
	}

	return config, nil
}
