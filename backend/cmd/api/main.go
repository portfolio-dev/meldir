package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"meldir-backend/internal/config"
	delivery "meldir-backend/internal/delivery/http"
	"meldir-backend/internal/delivery/http/handler"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("🏛️ Starting PT. Melayani Digital Raya (meldir-backend) v1.0.0...")
	log.Printf("📍 Environment: %s | Port: %s", cfg.Environment, cfg.Port)

	// Initialize Handlers
	healthHandler := handler.NewHealthHandler(cfg.Environment)

	// Setup Router
	router := delivery.NewRouter(healthHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", cfg.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("🚀 Server listening on http://127.0.0.1:%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server error: %v", err)
		}
	}()

	<-stop
	log.Println("🛑 Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("⚠️ Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited cleanly.")
}
