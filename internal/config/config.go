package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type (
	// Config содержит общую конфигурацию приложения.
	Config struct {
		// ServerConfig содержит конфигурацию сервера.
		Server ServerConfig `json:"server"`
	}
	// ServerConfig содержит конфигурацию сервера.
	ServerConfig struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}
)

func New() *Config {
	return &Config{
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: 8080,
		},
	}
}

// Load загружает и возвращает конфигурацию приложения.
func Load() (*Config, error) {
	cfg := New()
	path := parseOptions()

	if path != "" {
		if err := cfg.parseFile(path); err != nil {
			return nil, err
		}
	}

	cfg.overrideFromEnv()

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) parseFile(path string) error {
	cleanPath := filepath.Clean(path)
	file, err := os.Open(cleanPath)
	if err != nil {
		return fmt.Errorf("failed to open config: %w", err)
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(c)
}

func (c *Config) overrideFromEnv() {
	if host := os.Getenv("SERVER_HOST"); host != "" {
		c.Server.Host = host
	}

	if portStr := os.Getenv("SERVER_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			c.Server.Port = port
		}
	}
}
