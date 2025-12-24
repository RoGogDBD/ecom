package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RoGogDBD/ecom/internal/config"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "critical error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("could not load config: %w", err)
	}

	log.Printf("Starting server on %s:%d", cfg.Server.Host, cfg.Server.Port)

	return nil
}
