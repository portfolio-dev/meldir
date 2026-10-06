package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthHandler struct {
	env string
}

func NewHealthHandler(env string) *HealthHandler {
	return &HealthHandler{env: env}
}

type HealthResponse struct {
	Status      string    `json:"status"`
	Service     string    `json:"service"`
	Environment string    `json:"environment"`
	Host        string    `json:"host"`
	Timestamp   time.Time `json:"timestamp"`
	Version     string    `json:"version"`
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status:      "healthy",
		Service:     "meldir-enterprise-backend",
		Environment: h.env,
		Host:        r.Host,
		Timestamp:   time.Now().UTC(),
		Version:     "1.0.0-prod",
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
