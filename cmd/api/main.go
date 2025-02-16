package main

import (
	"context"
	"log"

	"os"
	"os/signal"
	"syscall"

	"github.com/epuerta9/openchef/internal/config"
	"github.com/epuerta9/openchef/internal/database"
	"github.com/epuerta9/openchef/internal/nats"
	"github.com/epuerta9/openchef/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize NATS
	ns, err := nats.New()
	if err != nil {
		log.Fatalf("Failed to start NATS: %v", err)
	}

	// Initialize database
	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create server instance
	srv := server.New(cfg, db)

	// Start server in a goroutine
	go func() {
		if err := srv.Start(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	// Shutdown everything gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	if err := ns.Shutdown(ctx); err != nil {
		log.Printf("NATS shutdown error: %v", err)
	}
}
