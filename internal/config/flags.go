package config

import (
	"flag"
	"os"
)

func parseOptions() string {
	var path string

	flag.StringVar(&path, "c", "", "Путь к конфигурационному файлу (сокращенный)")
	flag.StringVar(&path, "config", "", "Путь к конфигурационному файлу")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG")
	}

	return path
}
