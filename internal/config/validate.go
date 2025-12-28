package config

import "fmt"

// Validate проверяет корректность конфигурации.
func (c *Config) validate() error {
	if c == nil {
		return fmt.Errorf("config is nil")
	}

	if c.Server.Host == "" {
		return fmt.Errorf("server.host is required")
	}
	if c.Server.Port <= 0 {
		return fmt.Errorf("server.port must be > 0")
	}

	return nil
}
