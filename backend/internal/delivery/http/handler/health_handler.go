package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"meldir-backend/internal/infrastructure/cache"
	"meldir-backend/internal/infrastructure/database"
)

type HealthHandler struct {
	env   string
	db    *database.PostgresDB
	redis *cache.RedisClient
}

func NewHealthHandler(env string, db *database.PostgresDB, redis *cache.RedisClient) *HealthHandler {
	return &HealthHandler{
		env:   env,
		db:    db,
		redis: redis,
	}
}

type HealthResponse struct {
	Status      string                 `json:"status"`
	Service     string                 `json:"service"`
	Environment string                 `json:"environment"`
	Host        string                 `json:"host"`
	Timestamp   time.Time              `json:"timestamp"`
	Version     string                 `json:"version"`
	Databases   map[string]string      `json:"databases"`
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	dbStatus := "connected"
	if h.db != nil {
		if err := h.db.Pool.Ping(ctx); err != nil {
			dbStatus = "error: " + err.Error()
		}
	} else {
		dbStatus = "uninitialized"
	}

	redisStatus := "connected"
	if h.redis != nil {
		if err := h.redis.Client.Ping(ctx).Err(); err != nil {
			redisStatus = "error: " + err.Error()
		}
	} else {
		redisStatus = "uninitialized"
	}

	overallStatus := "healthy"
	if dbStatus != "connected" || redisStatus != "connected" {
		overallStatus = "degraded"
	}

	resp := HealthResponse{
		Status:      overallStatus,
		Service:     "meldir-enterprise-backend",
		Environment: h.env,
		Host:        r.Host,
		Timestamp:   time.Now().UTC(),
		Version:     "1.0.0-prod",
		Databases: map[string]string{
			"postgresql": dbStatus,
			"redis":      redisStatus,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *HealthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "pong",
		"server":  "PT. Melayani Digital Raya (meldir.id)",
		"time":    time.Now().UTC(),
	})
}
