package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"meldir-backend/internal/delivery/http/middleware"
	"meldir-backend/internal/domain"
	"meldir-backend/internal/usecase"
)

type AccountingHandler struct {
	accountingUsecase usecase.AccountingUsecase
}

func NewAccountingHandler(accountingUsecase usecase.AccountingUsecase) *AccountingHandler {
	return &AccountingHandler{accountingUsecase: accountingUsecase}
}

func (h *AccountingHandler) ListJournals(w http.ResponseWriter, r *http.Request) {
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

	limit := 100
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	journals, err := h.accountingUsecase.ListJournals(r.Context(), user, limit)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    journals,
	})
}

func (h *AccountingHandler) CreateJournal(w http.ResponseWriter, r *http.Request) {
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

	var req domain.CreateJournalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload JSON tidak valid",
		})
		return
	}

	entry, err := h.accountingUsecase.CreateJournal(r.Context(), user, req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Transaksi jurnal berhasil diposting ke Buku Besar secara seimbang",
		"data":    entry,
	})
}
