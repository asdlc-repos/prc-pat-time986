package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asdlc-repos/prc-pat-time986/time-service/handlers"
	"github.com/asdlc-repos/prc-pat-time986/time-service/middleware"
)

const defaultPort = "9090"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/time", handlers.TimeHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)

	loggedMux := middleware.Logger(mux)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      loggedMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in a goroutine so we can listen for OS signals concurrently.
	errCh := make(chan error, 1)
	go func() {
		fmt.Fprintf(os.Stdout, "time-service listening on :%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("listen: %w", err)
		}
	}()

	// Wait for SIGTERM or SIGINT for graceful shutdown.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sig := <-quit:
		fmt.Fprintf(os.Stdout, "received signal %s, shutting down gracefully\n", sig)
	case err := <-errCh:
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "graceful shutdown failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "server stopped")
}
