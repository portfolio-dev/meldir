package http

import (
	"net/http"
	"strings"

	"meldir-backend/internal/delivery/http/handler"
)

func NewRouter(healthHandler *handler.HealthHandler) http.Handler {
	mux := http.NewServeMux()

	// Health check endpoints
	mux.HandleFunc("/api/health", healthHandler.CheckHealth)
	mux.HandleFunc("/api/v1/ping", healthHandler.Ping)

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

		next.ServeHTTP(w)
	})
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log basic request information
		next.ServeHTTP(w, r)
	})
}
