package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config holds all configs.
type Config struct {
	GRPCServer GRPCConfig   `yaml:"gRPC_server"`
	Logger     LoggerConfig `yaml:"logger"`
}

// GRPCConfig configuration of grpc-instance service.
type GRPCConfig struct {
	Addr string `yaml:"addr"`
}

// LoggerConfig logger configuration.
//
// Filename - log file name.
//
// MaxSize - max log file size.
type LoggerConfig struct {
	Filename string `yaml:"file_name"`
	MaxSize  int    `yaml:"max_size"`
	Level    string `yaml:"level"`
}

// LoadFromFile create configuration from file.
func LoadFromFile(fileName string) (Config, error) {
	cfg := Config{}

	configBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(configBytes, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
