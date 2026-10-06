package http

import (
	"net/http"
	"strings"

	"meldir-backend/internal/delivery/http/handler"
	"meldir-backend/internal/delivery/http/middleware"
)

func NewRouter(
	healthHandler *handler.HealthHandler,
	authHandler *handler.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
) http.Handler {
	mux := http.NewServeMux()

	// 1. Health check & Ping endpoints
	mux.HandleFunc("/api/health", healthHandler.CheckHealth)
	mux.HandleFunc("/api/v1/ping", healthHandler.Ping)

	// 2. Auth Endpoints
	mux.HandleFunc("/api/v1/auth/login", authHandler.Login)
	mux.HandleFunc("/api/v1/user/profile", authMiddleware.Authenticate(authHandler.GetProfile))
	mux.HandleFunc("/api/v1/auth/logout", authMiddleware.Authenticate(authHandler.Logout))

	// Wrap with Global Middlewares (CORS, Logging, Recovery)
	return withCORS(withLogging(mux))
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		// Izinkan subdomain meldir.id dan localhost saat pengujian
		if origin != "" && (strings.HasSuffix(origin, "meldir.id") || strings.Contains(origin, "localhost")) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
