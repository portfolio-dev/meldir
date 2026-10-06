package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"meldir-backend/internal/delivery/http/middleware"
	"meldir-backend/internal/domain"
	"meldir-backend/internal/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	role := r.URL.Query().Get("role")
	search := r.URL.Query().Get("q")
	if search == "" {
		search = r.URL.Query().Get("search")
	}

	users, err := h.userUsecase.ListUsers(r.Context(), role, search)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    users,
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req domain.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	user, err := h.userUsecase.CreateUser(r.Context(), req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Akun pengguna berhasil dibuat",
		"data":    user,
	})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req domain.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	user, err := h.userUsecase.UpdateUser(r.Context(), req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Data akun pengguna berhasil diperbarui",
		"data":    user,
	})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete && r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	var id int64
	var err error

	if idStr != "" {
		id, err = strconv.ParseInt(idStr, 10, 64)
	} else {
		var body struct {
			ID int64 `json:"id"`
		}
		if decodeErr := json.NewDecoder(r.Body).Decode(&body); decodeErr == nil {
			id = body.ID
		}
	}

	if err != nil || id <= 0 {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "ID pengguna tidak valid",
		})
		return
	}

	var requesterID int64
	if reqVal := r.Context().Value(middleware.ContextKeyUserID); reqVal != nil {
		requesterID = reqVal.(int64)
	}

	if err := h.userUsecase.DeleteUser(r.Context(), id, requesterID); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Pengguna berhasil dihapus dari sistem",
	})
}
