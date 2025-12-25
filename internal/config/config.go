package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 8080

	envServerHost = "SERVER_HOST"
	envServerPort = "SERVER_PORT"
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

func NewLoad() *Config {
	return &Config{
		Server: ServerConfig{
			Host: defaultHost,
			Port: defaultPort,
		},
	}
}

// Load загружает и возвращает конфигурацию приложения.
func Load() (*Config, error) {
	cfg := NewLoad()
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

func (c *Config) parseFile(path string) (err error) {
	cleanPath := filepath.Clean(path)
	file, err := os.Open(cleanPath)
	if err != nil {
		return fmt.Errorf("failed to open config: %w", err)
	}

	// В данном случае вполне уместно игнорировать ошибку закрытия файла,
	// так как мы всё равно возвращаем ошибку из Decode. Как по мне - это холиварный вопрос.
	// Мою мысль хорошо описывает пост с Reddit: https://www.reddit.com/r/golang/comments/yrgths/comment/ivu2k9s/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button.
	// Но что бы литр не ругался, используем анонимную функцию для закрытия файла.
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()
	// Дефолтный подход - defer file.Close().

	err = json.NewDecoder(file).Decode(c)
	return
}

func (c *Config) overrideFromEnv() {
	if host := os.Getenv(envServerHost); host != "" {
		c.Server.Host = host
	}

	if portStr := os.Getenv(envServerPort); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			c.Server.Port = port
		}
	}
}
