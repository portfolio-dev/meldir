package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"meldir-backend/internal/delivery/http/middleware"
	"meldir-backend/internal/domain"
	"meldir-backend/internal/usecase"
)

type ChatHandler struct {
	chatUsecase usecase.ChatUsecase
}

func NewChatHandler(chatUsecase usecase.ChatUsecase) *ChatHandler {
	return &ChatHandler{chatUsecase: chatUsecase}
}

// StartPublicChat initiates a new session after visitor completes pre-chat form
func (h *ChatHandler) StartPublicChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req domain.StartChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format data onboarding chat tidak valid",
		})
		return
	}

	session, err := h.chatUsecase.StartChat(r.Context(), req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Sesi chat online berhasil dimulai",
		"data":    session,
	})
}

// SendVisitorMessage allows the public visitor to send messages
func (h *ChatHandler) SendVisitorMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req domain.SendChatMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format pesan tidak valid",
		})
		return
	}

	msg, err := h.chatUsecase.SendVisitorMessage(r.Context(), req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    msg,
	})
}

// GetPublicMessages polls or reads messages for a visitor session
func (h *ChatHandler) GetPublicMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	sessionCode := r.URL.Query().Get("session_code")
	if sessionCode == "" {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "session_code diperlukan",
		})
		return
	}

	session, msgs, err := h.chatUsecase.GetSessionMessagesByCode(r.Context(), sessionCode)
	if err != nil {
		writeJSONResponse(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success":  true,
		"session":  session,
		"messages": msgs,
	})
}

// ListOfficeSessions allows Office personnel to view all active and historical chat sessions
func (h *ChatHandler) ListOfficeSessions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	status := r.URL.Query().Get("status")
	sessions, err := h.chatUsecase.ListSessions(r.Context(), status)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    sessions,
	})
}

// GetOfficeSessionDetail returns full transcript of a session for Office staff
func (h *ChatHandler) GetOfficeSessionDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("session_id")
	sessionID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || sessionID <= 0 {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "session_id tidak valid",
		})
		return
	}

	session, msgs, err := h.chatUsecase.GetSessionMessagesByID(r.Context(), sessionID)
	if err != nil {
		writeJSONResponse(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success":  true,
		"session":  session,
		"messages": msgs,
	})
}

// SendOfficeReply allows staff to reply directly to visitor in real time
func (h *ChatHandler) SendOfficeReply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req domain.OfficeChatReplyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format payload pesan balasan tidak valid",
		})
		return
	}

	if req.AgentName == "" {
		if emailVal := r.Context().Value(middleware.ContextKeyEmail); emailVal != nil {
			if emailStr, ok := emailVal.(string); ok && emailStr != "" {
				req.AgentName = emailStr
			}
		}
		if req.AgentName == "" {
			req.AgentName = "Staff Officer Meldir"
		}
	}

	msg, err := h.chatUsecase.SendAgentReply(r.Context(), req)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Balasan berhasil dikirim ke pengunjung",
		"data":    msg,
	})
}

// UpdateOfficeSessionStatus updates session state (e.g. active, closed, follow_up)
func (h *ChatHandler) UpdateOfficeSessionStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPatch {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req domain.UpdateChatStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Format data status tidak valid",
		})
		return
	}

	if err := h.chatUsecase.UpdateStatus(r.Context(), req.SessionID, req.Status); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Status percakapan berhasil diperbarui",
	})
}
