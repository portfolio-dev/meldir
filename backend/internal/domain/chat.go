package domain

import "time"

type ChatSession struct {
	ID              int64         `json:"id"`
	SessionCode     string        `json:"session_code"`
	VisitorName     string        `json:"visitor_name"`
	VisitorPhone    string        `json:"visitor_phone"`
	VisitorEmail    string        `json:"visitor_email"`
	ServiceInterest string        `json:"service_interest"`
	InitialMessage  string        `json:"initial_message"`
	Status          string        `json:"status"` // 'active', 'closed', 'follow_up'
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Messages        []ChatMessage `json:"messages,omitempty"`
}

type ChatMessage struct {
	ID         int64     `json:"id"`
	SessionID  int64     `json:"session_id"`
	SenderType string    `json:"sender_type"` // 'visitor', 'agent', 'system'
	SenderName string    `json:"sender_name"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}

type StartChatRequest struct {
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	ServiceInterest string `json:"service_interest"`
	Message         string `json:"message"`
}

type SendChatMessageRequest struct {
	SessionCode string `json:"session_code"`
	SenderName  string `json:"sender_name"`
	Message     string `json:"message"`
}

type OfficeChatReplyRequest struct {
	SessionID int64  `json:"session_id"`
	AgentName string `json:"agent_name"`
	Message   string `json:"message"`
}

type UpdateChatStatusRequest struct {
	SessionID int64  `json:"session_id"`
	Status    string `json:"status"`
}
