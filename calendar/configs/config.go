package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config main application config.
type Config struct {
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
}

// ServerConfig http server config.
type ServerConfig struct {
	Addr string `yaml:"http_listen"`
}

// LoggerConfig zap logger config.
type LoggerConfig struct {
	LogFile  string `yaml:"log_file"`
	LogLevel string `yaml:"log_level"`
}

// LoadConfiguration unmarshal configuration from yaml file with filePath path
func LoadConfiguration(filePath string) (Config, error) {
	// read from file
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	// initialize configuration
	var conf Config
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
