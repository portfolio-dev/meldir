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
	"meldir-backend/internal/delivery/http/middleware"
	"meldir-backend/internal/infrastructure/cache"
	"meldir-backend/internal/infrastructure/database"
	"meldir-backend/internal/pkg/token"
	"meldir-backend/internal/repository"
	"meldir-backend/internal/usecase"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("🏛️ Starting PT. Melayani Digital Raya (meldir-backend) v1.0.0...")
	log.Printf("📍 Environment: %s | Port: %s", cfg.Environment, cfg.Port)

	// 1. Inisialisasi Database PostgreSQL
	db, dbErr := database.NewPostgresDB(cfg)
	if dbErr != nil {
		log.Printf("⚠️ Gagal inisialisasi PostgreSQL pool: %v", dbErr)
	} else {
		defer db.Close()
	}

	// 2. Inisialisasi Cache & Blacklist Redis
	redisClient, redisErr := cache.NewRedisClient(cfg)
	if redisErr != nil {
		log.Printf("⚠️ Gagal inisialisasi Redis client: %v", redisErr)
	} else {
		defer redisClient.Close()
	}

	// 3. Inisialisasi Security & JWT Manager (24 Jam)
	jwtManager := token.NewJWTManager(cfg.JWTSecret, 24*time.Hour)

	// 4. Inisialisasi Repository & Usecase
	var userRepo repository.UserRepository
	if db != nil {
		userRepo = repository.NewUserRepository(db.Pool)
	}
	authUsecase := usecase.NewAuthUsecase(userRepo, jwtManager, redisClient)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// 5. Inisialisasi Handlers & Middleware
	healthHandler := handler.NewHealthHandler(cfg.Environment, db, dbErr, redisClient, redisErr)
	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	authMiddleware := middleware.NewAuthMiddleware(jwtManager, redisClient)

	// 6. Setup Router
	router := delivery.NewRouter(healthHandler, authHandler, userHandler, authMiddleware)

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
		log.Printf("🚀 Server RESTful API listening on http://127.0.0.1:%s", cfg.Port)
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
