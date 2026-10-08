package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"meldir-backend/internal/domain"
)

type ChatRepository interface {
	CreateSession(ctx context.Context, session *domain.ChatSession) (*domain.ChatSession, error)
	GetSessionByCode(ctx context.Context, code string) (*domain.ChatSession, error)
	GetSessionByID(ctx context.Context, id int64) (*domain.ChatSession, error)
	ListSessions(ctx context.Context, status string) ([]domain.ChatSession, error)
	AddMessage(ctx context.Context, msg *domain.ChatMessage) (*domain.ChatMessage, error)
	GetMessages(ctx context.Context, sessionID int64) ([]domain.ChatMessage, error)
	UpdateSessionStatus(ctx context.Context, sessionID int64, status string) error
}

type chatRepo struct {
	pool *pgxpool.Pool
}

func NewChatRepository(pool *pgxpool.Pool) ChatRepository {
	return &chatRepo{pool: pool}
}

func (r *chatRepo) CreateSession(ctx context.Context, session *domain.ChatSession) (*domain.ChatSession, error) {
	now := time.Now().UTC()
	session.CreatedAt = now
	session.UpdatedAt = now
	if session.Status == "" {
		session.Status = "active"
	}
	if session.SessionCode == "" {
		session.SessionCode = fmt.Sprintf("CHAT-%d-%d", now.Unix(), time.Now().Nanosecond()%10000)
	}

	query := `
		INSERT INTO public_chat_sessions (
			session_code, visitor_name, visitor_phone, visitor_email,
			service_interest, initial_message, status, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		) RETURNING id
	`

	err := r.pool.QueryRow(
		ctx, query,
		session.SessionCode, session.VisitorName, session.VisitorPhone, session.VisitorEmail,
		session.ServiceInterest, session.InitialMessage, session.Status, session.CreatedAt, session.UpdatedAt,
	).Scan(&session.ID)

	if err != nil {
		log.Printf("❌ ChatRepo.CreateSession error: %v", err)
		return nil, err
	}

	return session, nil
}

func (r *chatRepo) GetSessionByCode(ctx context.Context, code string) (*domain.ChatSession, error) {
	query := `
		SELECT id, session_code, visitor_name, visitor_phone, visitor_email,
		       service_interest, initial_message, status, created_at, updated_at
		FROM public_chat_sessions
		WHERE session_code = $1
	`
	var s domain.ChatSession
	err := r.pool.QueryRow(ctx, query, code).Scan(
		&s.ID, &s.SessionCode, &s.VisitorName, &s.VisitorPhone, &s.VisitorEmail,
		&s.ServiceInterest, &s.InitialMessage, &s.Status, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *chatRepo) GetSessionByID(ctx context.Context, id int64) (*domain.ChatSession, error) {
	query := `
		SELECT id, session_code, visitor_name, visitor_phone, visitor_email,
		       service_interest, initial_message, status, created_at, updated_at
		FROM public_chat_sessions
		WHERE id = $1
	`
	var s domain.ChatSession
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&s.ID, &s.SessionCode, &s.VisitorName, &s.VisitorPhone, &s.VisitorEmail,
		&s.ServiceInterest, &s.InitialMessage, &s.Status, &s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *chatRepo) ListSessions(ctx context.Context, status string) ([]domain.ChatSession, error) {
	query := `
		SELECT id, session_code, visitor_name, visitor_phone, visitor_email,
		       service_interest, initial_message, status, created_at, updated_at
		FROM public_chat_sessions
		WHERE ($1 = '' OR status = $1)
		ORDER BY updated_at DESC, created_at DESC
	`
	rows, err := r.pool.Query(ctx, query, status)
	if err != nil {
		log.Printf("❌ ChatRepo.ListSessions query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	sessions := make([]domain.ChatSession, 0)
	for rows.Next() {
		var s domain.ChatSession
		err := rows.Scan(
			&s.ID, &s.SessionCode, &s.VisitorName, &s.VisitorPhone, &s.VisitorEmail,
			&s.ServiceInterest, &s.InitialMessage, &s.Status, &s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			log.Printf("❌ ChatRepo.ListSessions scan error: %v", err)
			return nil, err
		}
		sessions = append(sessions, s)
	}

	return sessions, nil
}

func (r *chatRepo) AddMessage(ctx context.Context, msg *domain.ChatMessage) (*domain.ChatMessage, error) {
	now := time.Now().UTC()
	msg.CreatedAt = now

	query := `
		INSERT INTO public_chat_messages (
			session_id, sender_type, sender_name, message, created_at
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING id
	`

	err := r.pool.QueryRow(
		ctx, query,
		msg.SessionID, msg.SenderType, msg.SenderName, msg.Message, msg.CreatedAt,
	).Scan(&msg.ID)

	if err != nil {
		log.Printf("❌ ChatRepo.AddMessage insert error: %v", err)
		return nil, err
	}

	// Update session updated_at
	_, _ = r.pool.Exec(ctx, "UPDATE public_chat_sessions SET updated_at = $1 WHERE id = $2", now, msg.SessionID)

	return msg, nil
}

func (r *chatRepo) GetMessages(ctx context.Context, sessionID int64) ([]domain.ChatMessage, error) {
	query := `
		SELECT id, session_id, sender_type, sender_name, message, created_at
		FROM public_chat_messages
		WHERE session_id = $1
		ORDER BY created_at ASC
	`
	rows, err := r.pool.Query(ctx, query, sessionID)
	if err != nil {
		log.Printf("❌ ChatRepo.GetMessages query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	messages := make([]domain.ChatMessage, 0)
	for rows.Next() {
		var m domain.ChatMessage
		err := rows.Scan(&m.ID, &m.SessionID, &m.SenderType, &m.SenderName, &m.Message, &m.CreatedAt)
		if err != nil {
			log.Printf("❌ ChatRepo.GetMessages scan error: %v", err)
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func (r *chatRepo) UpdateSessionStatus(ctx context.Context, sessionID int64, status string) error {
	now := time.Now().UTC()
	query := `UPDATE public_chat_sessions SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.pool.Exec(ctx, query, status, now, sessionID)
	return err
}
