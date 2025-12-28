package config

import (
	"testing"
)

func TestConfig_validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name:    "nil config",
			config:  nil,
			wantErr: true,
		},
		{
			name: "пустой host",
			config: &Config{
				Server: ServerConfig{
					Host: "",
					Port: 8080,
				},
			},
			wantErr: true,
		},
		{
			name: "нулевой порт",
			config: &Config{
				Server: ServerConfig{
					Host: "localhost",
					Port: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "<0 порт",
			config: &Config{
				Server: ServerConfig{
					Host: "localhost",
					Port: -1,
				},
			},
			wantErr: true,
		},
		{
			name: "валидный конфиг",
			config: &Config{
				Server: ServerConfig{
					Host: "localhost",
					Port: 8080,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() ошибка = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
