package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config holds all configs.
type Config struct {
	Stan   StanConfig   `yaml:"stan"`
	Logger LoggerConfig `yaml:"logger"`
}

// StanConfig - configuration to connect to nats-streaming (stan)
type StanConfig struct {
	ClusterName string `yaml:"cluster_name"`
	ClientID    string `yaml:"client_id"`
	Addr        string `yaml:"addr"`
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
