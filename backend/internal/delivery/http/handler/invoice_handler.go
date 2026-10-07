package handler

import (
	"encoding/json"
	"net/http"

	"meldir-backend/internal/delivery/http/middleware"
	"meldir-backend/internal/domain"
	"meldir-backend/internal/usecase"
)

type InvoiceHandler struct {
	invoiceUsecase usecase.InvoiceUsecase
}

func NewInvoiceHandler(invoiceUsecase usecase.InvoiceUsecase) *InvoiceHandler {
	return &InvoiceHandler{invoiceUsecase: invoiceUsecase}
}

func (h *InvoiceHandler) ListInvoices(w http.ResponseWriter, r *http.Request) {
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

	invoices, err := h.invoiceUsecase.ListInvoices(r.Context(), user)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    invoices,
	})
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
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

	var req domain.CreateInvoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	inv, err := h.invoiceUsecase.CreateInvoice(r.Context(), user, req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Faktur invoice resmi berhasil diterbitkan",
		"data":    inv,
	})
}

func (h *InvoiceHandler) UpdateInvoiceStatus(w http.ResponseWriter, r *http.Request) {
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

	var req domain.UpdateInvoiceStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	if err := h.invoiceUsecase.UpdateStatus(r.Context(), user, req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Status pembayaran faktur berhasil diperbarui",
	})
}
