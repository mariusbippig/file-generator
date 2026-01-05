package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Filename      string
	FileType      string
	FileSizeBytes int
	FilesAmount   int
	Logger        Logger
}

type Logger struct {
	Driver   string
	LogLevel string
}

// NewConfig inits a new config by configured config source.
func NewConfig(path string) (Config, error) {
	// TODO add new config origins as drivers here.
	driver := "env"

	switch driver {
	default:
		return NewConfigFromEnv(path)
	}
}

func NewConfigFromEnv(path string) (Config, error) {
	viper.SetConfigName(path)
	viper.AddConfigPath(".")

	var fileLookupError viper.ConfigFileNotFoundError
	err := viper.ReadInConfig()

	if !errors.As(err, &fileLookupError) && err != nil {
		return Config{}, err
	}

	viper.SetDefault("filename", "foobar")
	viper.SetDefault("fileType", "txt")
	viper.SetDefault("size", "size")
	viper.SetDefault("amount", "size")
	viper.SetDefault("logger.driver", "console")
	viper.SetDefault("logger.logLevel", "debug")

	viper.BindEnv("filename", "FILENAME")
	viper.BindEnv("fileType", "FILETYPE")
	viper.BindEnv("size", "SIZE")
	viper.BindEnv("size", "SIZE")
	viper.BindEnv("logger.driver", "LOGGER")
	viper.BindEnv("logger.logLevel", "LOG_LEVEL")

	config := Config{
		Filename:      viper.GetString("filename"),
		FileType:      viper.GetString("fileType"),
		FileSizeBytes: viper.GetInt("size"),
		FilesAmount:   viper.GetInt("amount"),
		Logger: Logger{
			Driver:   viper.GetString("logger.driver"),
			LogLevel: viper.GetString("logger.logLevel"),
		},
	}

	return config, nil
}
