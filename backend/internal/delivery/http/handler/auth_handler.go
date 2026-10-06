package handler

import (
	"encoding/json"
	"net/http"

	"meldir-backend/internal/delivery/http/middleware"
	"meldir-backend/internal/domain"
	"meldir-backend/internal/usecase"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Payload JSON tidak valid",
		})
		return
	}

	resp, err := h.authUsecase.Login(r.Context(), req)
	if err != nil {
		writeJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Login berhasil",
		"data":    resp,
	})
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	userIDVal := r.Context().Value(middleware.ContextKeyUserID)
	if userIDVal == nil {
		writeJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"error":   "Sesi pengguna tidak valid",
		})
		return
	}

	userID := userIDVal.(int64)
	user, err := h.authUsecase.GetProfile(r.Context(), userID)
	if err != nil {
		writeJSONResponse(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	tokenVal := r.Context().Value(middleware.ContextKeyToken)
	var rawToken string
	if tokenVal != nil {
		rawToken = tokenVal.(string)
	}

	_ = h.authUsecase.Logout(r.Context(), rawToken)

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Berhasil logout, sesi token telah dicabut",
	})
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
