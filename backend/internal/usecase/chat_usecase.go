package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"meldir-backend/internal/domain"
	"meldir-backend/internal/repository"
)

type ChatUsecase interface {
	StartChat(ctx context.Context, req domain.StartChatRequest) (*domain.ChatSession, error)
	SendVisitorMessage(ctx context.Context, req domain.SendChatMessageRequest) (*domain.ChatMessage, error)
	GetSessionMessagesByCode(ctx context.Context, code string) (*domain.ChatSession, []domain.ChatMessage, error)
	ListSessions(ctx context.Context, status string) ([]domain.ChatSession, error)
	GetSessionMessagesByID(ctx context.Context, sessionID int64) (*domain.ChatSession, []domain.ChatMessage, error)
	SendAgentReply(ctx context.Context, req domain.OfficeChatReplyRequest) (*domain.ChatMessage, error)
	UpdateStatus(ctx context.Context, sessionID int64, status string) error
}

type chatUsecase struct {
	chatRepo repository.ChatRepository
	leadRepo repository.LeadRepository
}

func NewChatUsecase(chatRepo repository.ChatRepository, leadRepo repository.LeadRepository) ChatUsecase {
	return &chatUsecase{
		chatRepo: chatRepo,
		leadRepo: leadRepo,
	}
}

func (u *chatUsecase) StartChat(ctx context.Context, req domain.StartChatRequest) (*domain.ChatSession, error) {
	if u.chatRepo == nil {
		return nil, errors.New("koneksi repositori chat belum siap")
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("nama lengkap wajib diisi")
	}
	phone := strings.TrimSpace(req.Phone)
	if phone == "" {
		return nil, errors.New("nomor kontak WhatsApp atau HP wajib diisi")
	}
	msgText := strings.TrimSpace(req.Message)
	if msgText == "" {
		return nil, errors.New("pesan atau pertanyaan awal wajib diisi")
	}

	interest := strings.TrimSpace(req.ServiceInterest)
	if interest == "" {
		interest = "Konsultasi Umum"
	}

	session := &domain.ChatSession{
		VisitorName:     name,
		VisitorPhone:    phone,
		VisitorEmail:    strings.TrimSpace(req.Email),
		ServiceInterest: interest,
		InitialMessage:  msgText,
		Status:          "active",
	}

	createdSession, err := u.chatRepo.CreateSession(ctx, session)
	if err != nil {
		return nil, fmt.Errorf("gagal memulai sesi chat: %w", err)
	}

	// 1. Catat pesan awal pengunjung
	_, _ = u.chatRepo.AddMessage(ctx, &domain.ChatMessage{
		SessionID:  createdSession.ID,
		SenderType: "visitor",
		SenderName: name,
		Message:    msgText,
	})

	// 2. Kirim pesan sambutan otomatis dari AI / Konsultan Meldir
	welcomeMsg := fmt.Sprintf("Halo Bpk/Ibu %s! Terima kasih telah menghubungi PT. Melayani Digital Raya. Konsultan kami telah menerima pertanyaan Anda mengenai '%s' dan siap membantu Anda di sini.", name, interest)
	_, _ = u.chatRepo.AddMessage(ctx, &domain.ChatMessage{
		SessionID:  createdSession.ID,
		SenderType: "agent",
		SenderName: "Meldir Desk Officer",
		Message:    welcomeMsg,
	})

	// 3. Rekam ke Inbound Leads Pipeline jika leadRepo tersedia
	if u.leadRepo != nil {
		_, _ = u.leadRepo.CreateLead(ctx, &domain.Lead{
			Name:            name,
			Email:           strings.TrimSpace(req.Email),
			PhoneWA:         phone,
			ServiceInterest: interest,
			Message:         fmt.Sprintf("[Live Chat - %s] %s", createdSession.SessionCode, msgText),
			Source:          "live_chat_web",
			Status:          domain.LeadStatusNew,
		})
	}

	// Ambil kembali seluruh pesan
	messages, err := u.chatRepo.GetMessages(ctx, createdSession.ID)
	if err == nil {
		createdSession.Messages = messages
	}

	return createdSession, nil
}

func (u *chatUsecase) SendVisitorMessage(ctx context.Context, req domain.SendChatMessageRequest) (*domain.ChatMessage, error) {
	if u.chatRepo == nil {
		return nil, errors.New("koneksi repositori chat belum siap")
	}

	code := strings.TrimSpace(req.SessionCode)
	if code == "" {
		return nil, errors.New("kode sesi chat tidak valid")
	}
	text := strings.TrimSpace(req.Message)
	if text == "" {
		return nil, errors.New("pesan tidak boleh kosong")
	}

	session, err := u.chatRepo.GetSessionByCode(ctx, code)
	if err != nil {
		return nil, errors.New("sesi percakapan tidak ditemukan")
	}

	senderName := strings.TrimSpace(req.SenderName)
	if senderName == "" {
		senderName = session.VisitorName
	}

	// Reopen session if closed
	if session.Status == "closed" {
		_ = u.chatRepo.UpdateSessionStatus(ctx, session.ID, "active")
	}

	return u.chatRepo.AddMessage(ctx, &domain.ChatMessage{
		SessionID:  session.ID,
		SenderType: "visitor",
		SenderName: senderName,
		Message:    text,
	})
}

func (u *chatUsecase) GetSessionMessagesByCode(ctx context.Context, code string) (*domain.ChatSession, []domain.ChatMessage, error) {
	if u.chatRepo == nil {
		return nil, nil, errors.New("koneksi repositori chat belum siap")
	}
	session, err := u.chatRepo.GetSessionByCode(ctx, strings.TrimSpace(code))
	if err != nil {
		return nil, nil, errors.New("sesi chat tidak ditemukan")
	}
	messages, err := u.chatRepo.GetMessages(ctx, session.ID)
	return session, messages, err
}

func (u *chatUsecase) ListSessions(ctx context.Context, status string) ([]domain.ChatSession, error) {
	if u.chatRepo == nil {
		return nil, errors.New("koneksi repositori chat belum siap")
	}
	return u.chatRepo.ListSessions(ctx, strings.TrimSpace(status))
}

func (u *chatUsecase) GetSessionMessagesByID(ctx context.Context, sessionID int64) (*domain.ChatSession, []domain.ChatMessage, error) {
	if u.chatRepo == nil {
		return nil, nil, errors.New("koneksi repositori chat belum siap")
	}
	session, err := u.chatRepo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return nil, nil, errors.New("sesi chat tidak ditemukan")
	}
	messages, err := u.chatRepo.GetMessages(ctx, sessionID)
	return session, messages, err
}

func (u *chatUsecase) SendAgentReply(ctx context.Context, req domain.OfficeChatReplyRequest) (*domain.ChatMessage, error) {
	if u.chatRepo == nil {
		return nil, errors.New("koneksi repositori chat belum siap")
	}
	if req.SessionID <= 0 {
		return nil, errors.New("ID sesi tidak valid")
	}
	text := strings.TrimSpace(req.Message)
	if text == "" {
		return nil, errors.New("isi pesan balasan tidak boleh kosong")
	}
	agent := strings.TrimSpace(req.AgentName)
	if agent == "" {
		agent = "Office Support Meldir"
	}

	return u.chatRepo.AddMessage(ctx, &domain.ChatMessage{
		SessionID:  req.SessionID,
		SenderType: "agent",
		SenderName: agent,
		Message:    text,
	})
}

func (u *chatUsecase) UpdateStatus(ctx context.Context, sessionID int64, status string) error {
	if u.chatRepo == nil {
		return errors.New("koneksi repositori chat belum siap")
	}
	if sessionID <= 0 {
		return errors.New("ID sesi tidak valid")
	}
	st := strings.TrimSpace(status)
	if st == "" {
		st = "active"
	}
	return u.chatRepo.UpdateSessionStatus(ctx, sessionID, st)
}
