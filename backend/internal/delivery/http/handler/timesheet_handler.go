package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"meldir-backend/internal/delivery/http/middleware"
	"meldir-backend/internal/domain"
	"meldir-backend/internal/usecase"
)

type TimesheetHandler struct {
	timesheetUsecase usecase.TimesheetUsecase
}

func NewTimesheetHandler(timesheetUsecase usecase.TimesheetUsecase) *TimesheetHandler {
	return &TimesheetHandler{timesheetUsecase: timesheetUsecase}
}

func (h *TimesheetHandler) ListTimesheets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	user := middleware.GetUserClaims(r)
	if user == nil || user.ID <= 0 {
		writeJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"error":   "Sesi pengguna tidak valid",
		})
		return
	}

	limit := 50
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	logs, err := h.timesheetUsecase.ListTimesheets(r.Context(), user, limit)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    logs,
	})
}

func (h *TimesheetHandler) LogWork(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	user := middleware.GetUserClaims(r)
	if user == nil || user.ID <= 0 {
		writeJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"error":   "Sesi pengguna tidak valid",
		})
		return
	}

	var req domain.CreateTimesheetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	logEntry, err := h.timesheetUsecase.LogWork(r.Context(), user, req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Waktu pengerjaan berhasil dicatat ke timesheet",
		"data":    logEntry,
	})
}

func (h *TimesheetHandler) GetSummary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	user := middleware.GetUserClaims(r)
	if user == nil || user.ID <= 0 {
		writeJSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"error":   "Sesi pengguna tidak valid",
		})
		return
	}

	summary, err := h.timesheetUsecase.GetSummary(r.Context(), user)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    summary,
	})
}
