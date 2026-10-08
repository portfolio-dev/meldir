package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"meldir-backend/internal/infrastructure/cache"
	"meldir-backend/internal/infrastructure/database"
)

type HealthHandler struct {
	env      string
	db       *database.PostgresDB
	dbErr    error
	redis    *cache.RedisClient
	redisErr error
}

func NewHealthHandler(env string, db *database.PostgresDB, dbErr error, redis *cache.RedisClient, redisErr error) *HealthHandler {
	return &HealthHandler{
		env:      env,
		db:       db,
		dbErr:    dbErr,
		redis:    redis,
		redisErr: redisErr,
	}
}

type HealthResponse struct {
	Status      string            `json:"status"`
	Service     string            `json:"service"`
	Environment string            `json:"environment"`
	Host        string            `json:"host"`
	Timestamp   time.Time         `json:"timestamp"`
	Version     string            `json:"version"`
	Databases   map[string]string `json:"databases"`
	Diagnostics map[string]string `json:"diagnostics,omitempty"`
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	dbStatus := "connected"
	if h.db != nil {
		if err := h.db.Pool.Ping(ctx); err != nil {
			dbStatus = "error: " + err.Error()
		}
	} else if h.dbErr != nil {
		dbStatus = "error: " + h.dbErr.Error()
	} else {
		dbStatus = "uninitialized (db is nil)"
	}

	redisStatus := "connected"
	if h.redis != nil {
		if err := h.redis.Client.Ping(ctx).Err(); err != nil {
			redisStatus = "error: " + err.Error()
		}
	} else if h.redisErr != nil {
		redisStatus = "error: " + h.redisErr.Error()
	} else {
		redisStatus = "uninitialized (redis is nil)"
	}

	diagnostics := make(map[string]string)
	if h.db != nil {
		var invCount int
		if err := h.db.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM invoices").Scan(&invCount); err != nil {
			diagnostics["invoices_table"] = "error: " + err.Error()
		} else {
			diagnostics["invoices_table"] = fmt.Sprintf("ok (%d rows)", invCount)
		}

		var jrnCount int
		if err := h.db.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM accounting_journals").Scan(&jrnCount); err != nil {
			diagnostics["journals_table"] = "error: " + err.Error()
		} else {
			diagnostics["journals_table"] = fmt.Sprintf("ok (%d rows)", jrnCount)
		}
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
		Diagnostics: diagnostics,
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
