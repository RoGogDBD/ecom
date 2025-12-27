package config

import (
	"os"
	"testing"
)

func TestNewDefault(t *testing.T) {
	cfg := NewDefault()

	if cfg == nil {
		t.Fatal("NewDefault() returned nil")
	}

	if cfg.Server.Host != defaultHost {
		t.Errorf("expected host %q, got %q", defaultHost, cfg.Server.Host)
	}

	if cfg.Server.Port != defaultPort {
		t.Errorf("expected port %d, got %d", defaultPort, cfg.Server.Port)
	}
}

func TestConfig_overrideFromEnv(t *testing.T) {
	tests := []struct {
		name     string
		envHost  string
		envPort  string
		wantHost string
		wantPort int
		wantErr  bool
	}{
		{
			name:     "no env vars",
			envHost:  "",
			envPort:  "",
			wantHost: defaultHost,
			wantPort: defaultPort,
			wantErr:  false,
		},
		{
			name:     "valid env vars",
			envHost:  "127.0.0.1",
			envPort:  "9090",
			wantHost: "127.0.0.1",
			wantPort: 9090,
			wantErr:  false,
		},
		{
			name:     "invalid port",
			envHost:  "127.0.0.1",
			envPort:  "invalid",
			wantHost: "127.0.0.1",
			wantPort: defaultPort,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origHost := os.Getenv(envServerHost)
			origPort := os.Getenv(envServerPort)
			defer func() {
				_ = os.Setenv(envServerHost, origHost)
				_ = os.Setenv(envServerPort, origPort)
			}()

			if tt.envHost != "" {
				_ = os.Setenv(envServerHost, tt.envHost)
			} else {
				_ = os.Unsetenv(envServerHost)
			}
			if tt.envPort != "" {
				_ = os.Setenv(envServerPort, tt.envPort)
			} else {
				_ = os.Unsetenv(envServerPort)
			}

			cfg := NewDefault()
			err := cfg.overrideFromEnv()

			if (err != nil) != tt.wantErr {
				t.Errorf("overrideFromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if cfg.Server.Host != tt.wantHost {
					t.Errorf("expected host %q, got %q", tt.wantHost, cfg.Server.Host)
				}
				if cfg.Server.Port != tt.wantPort {
					t.Errorf("expected port %d, got %d", tt.wantPort, cfg.Server.Port)
				}
			}
		})
	}
}
