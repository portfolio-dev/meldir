package handler

import (
	"encoding/json"
	"net/http"

	"meldir-backend/internal/delivery/http/middleware"
	"meldir-backend/internal/domain"
	"meldir-backend/internal/usecase"
)

type TicketHandler struct {
	ticketUsecase usecase.TicketUsecase
}

func NewTicketHandler(ticketUsecase usecase.TicketUsecase) *TicketHandler {
	return &TicketHandler{ticketUsecase: ticketUsecase}
}

func (h *TicketHandler) ListTickets(w http.ResponseWriter, r *http.Request) {
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

	filter := domain.TicketFilter{
		Status:   r.URL.Query().Get("status"),
		Priority: r.URL.Query().Get("priority"),
		Search:   r.URL.Query().Get("q"),
	}
	if filter.Search == "" {
		filter.Search = r.URL.Query().Get("search")
	}

	tickets, err := h.ticketUsecase.ListTickets(r.Context(), user, filter)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    tickets,
	})
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
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

	var req domain.CreateTicketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	ticket, err := h.ticketUsecase.CreateTicket(r.Context(), user, req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Tiket bantuan SLA berhasil diajukan",
		"data":    ticket,
	})
}

func (h *TicketHandler) UpdateTicketStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch && r.Method != http.MethodPost {
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

	var req domain.UpdateTicketStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	if err := h.ticketUsecase.UpdateStatus(r.Context(), user, req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Status tiket SLA berhasil diperbarui",
	})
}
