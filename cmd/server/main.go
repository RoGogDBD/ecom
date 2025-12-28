package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RoGogDBD/ecom/internal/config"
	"github.com/RoGogDBD/ecom/internal/handler"
	"github.com/RoGogDBD/ecom/internal/repository"
	"github.com/RoGogDBD/ecom/internal/service"
)

const (
	errServerShutdown = "server shutdown failed"
	errLoadConfig     = "could not load config"

	logServerStart = "Starting server on %s"
	logServerStop  = "Server stopped"
	logShutdown    = "Shutting down gracefully..."
	logHTTPError   = "HTTP server error: %v"
	logCriticalErr = "critical error: %v"

	shutdownTimeout = 10 * time.Second
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, logCriticalErr, err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("%s: %w", errLoadConfig, err)
	}

	storage := repository.NewTodoStorage()
	todoService := service.NewTodoService(storage)
	router := handler.NewRouter(todoService)
	httpHandler := handler.Conveyor(
		router,
		handler.LoggingMiddleware(log.Default()),
	)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: httpHandler,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf(logServerStart, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf(logHTTPError, err)
		}
	}()

	<-sigChan
	log.Println(logShutdown)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: %w", errServerShutdown, err)
	}

	log.Println(logServerStop)
	return nil
}
