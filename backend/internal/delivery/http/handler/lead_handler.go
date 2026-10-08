package handler

import (
	"encoding/json"
	"net/http"

	"meldir-backend/internal/domain"
	"meldir-backend/internal/usecase"
)

type LeadHandler struct {
	leadUsecase usecase.LeadUsecase
}

func NewLeadHandler(leadUsecase usecase.LeadUsecase) *LeadHandler {
	return &LeadHandler{leadUsecase: leadUsecase}
}

// CreatePublicLead can be called from public website without JWT authentication
func (h *LeadHandler) CreatePublicLead(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req domain.CreateLeadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	lead, err := h.leadUsecase.CreatePublicLead(r.Context(), req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success":   true,
		"message":   "Permintaan audit & konsultasi Anda telah berhasil dikirimkan.",
		"lead_code": lead.LeadCode,
		"data":      lead,
	})
}

// ListLeads can only be accessed by authorized staff (direktur, admin)
func (h *LeadHandler) ListLeads(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	statusFilter := r.URL.Query().Get("status")
	leads, err := h.leadUsecase.ListLeads(r.Context(), statusFilter)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    leads,
	})
}

// UpdateLeadStatus updates status and notes of a lead
func (h *LeadHandler) UpdateLeadStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPatch {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req domain.UpdateLeadStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	if err := h.leadUsecase.UpdateStatus(r.Context(), req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Status prospek berhasil diperbarui",
	})
}
