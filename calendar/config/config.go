package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config holds all configs.
type Config struct {
	GRPCServer GRPCConfig   `yaml:"gRPC_server"`
	Db         DbConfig     `yaml:"db"`
	Logger     LoggerConfig `yaml:"logger"`
}

// GRPCConfig configuration of grpc-instance service.
type GRPCConfig struct {
	Addr string `yaml:"addr"`
}

// DbConfig configuration of postgres database.
type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"ssl_mode"`
	MaxConn  int    `yaml:"max_conn"`
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
