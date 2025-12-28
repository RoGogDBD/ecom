package config

import (
	"flag"
	"os"
	"testing"
)

func TestParseOptions(t *testing.T) {
	originalCommandLine := flag.CommandLine
	originalArgs := os.Args
	originalEnv := os.Getenv("CONFIG")

	defer func() {
		flag.CommandLine = originalCommandLine
		os.Args = originalArgs
		if originalEnv == "" {
			_ = os.Unsetenv("CONFIG")
		} else {
			_ = os.Setenv("CONFIG", originalEnv)
		}
	}()

	tests := []struct {
		name     string
		args     []string
		envValue string
		want     string
	}{
		{
			name:     "по умолчанию пусто",
			args:     []string{"cmd"},
			envValue: "",
			want:     "",
		},
		{
			name:     "конфиг из env",
			args:     []string{"cmd"},
			envValue: "/tmp/config.json",
			want:     "/tmp/config.json",
		},
		{
			name:     "короткий флаг",
			args:     []string{"cmd", "-c", "/etc/app.json"},
			envValue: "/tmp/config.json",
			want:     "/etc/app.json",
		},
		{
			name:     "длинный флаг",
			args:     []string{"cmd", "-config", "/etc/app.json"},
			envValue: "",
			want:     "/etc/app.json",
		},
		{
			name:     "флаг перекрывает env",
			args:     []string{"cmd", "-config", "/etc/app.json"},
			envValue: "/tmp/config.json",
			want:     "/etc/app.json",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(tt.args[0], flag.ContinueOnError)
			os.Args = tt.args

			if tt.envValue == "" {
				_ = os.Unsetenv("CONFIG")
			} else {
				_ = os.Setenv("CONFIG", tt.envValue)
			}

			got := parseOptions()
			if got != tt.want {
				t.Fatalf("ожидалось %q, получено %q", tt.want, got)
			}
		})
	}
}
