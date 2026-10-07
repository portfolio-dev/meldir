package usecase

import (
	"context"
	"errors"
	"strings"

	"meldir-backend/internal/domain"
	"meldir-backend/internal/repository"
)

type TicketUsecase interface {
	ListTickets(ctx context.Context, currentUser *domain.User, filter domain.TicketFilter) ([]domain.Ticket, error)
	GetTicketByID(ctx context.Context, currentUser *domain.User, id int64) (*domain.Ticket, error)
	CreateTicket(ctx context.Context, currentUser *domain.User, req domain.CreateTicketRequest) (*domain.Ticket, error)
	UpdateStatus(ctx context.Context, currentUser *domain.User, req domain.UpdateTicketStatusRequest) error
}

type ticketUsecase struct {
	ticketRepo repository.TicketRepository
}

func NewTicketUsecase(ticketRepo repository.TicketRepository) TicketUsecase {
	return &ticketUsecase{ticketRepo: ticketRepo}
}

func (u *ticketUsecase) ListTickets(ctx context.Context, currentUser *domain.User, filter domain.TicketFilter) ([]domain.Ticket, error) {
	if u.ticketRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	// Jika pengguna adalah Klien, batasi hanya melihat tiket miliknya sendiri
	if currentUser.Role == domain.RoleKlien {
		filter.ClientID = currentUser.ID
	}

	return u.ticketRepo.ListTickets(ctx, filter)
}

func (u *ticketUsecase) GetTicketByID(ctx context.Context, currentUser *domain.User, id int64) (*domain.Ticket, error) {
	if u.ticketRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	ticket, err := u.ticketRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Klien hanya boleh membuka detail tiketnya sendiri
	if currentUser.Role == domain.RoleKlien && ticket.ClientID != currentUser.ID {
		return nil, errors.New("Anda tidak memiliki hak akses untuk tiket ini")
	}

	return ticket, nil
}

func (u *ticketUsecase) CreateTicket(ctx context.Context, currentUser *domain.User, req domain.CreateTicketRequest) (*domain.Ticket, error) {
	if u.ticketRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	if strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("judul kendala / tiket wajib diisi")
	}
	if strings.TrimSpace(req.Description) == "" {
		return nil, errors.New("deskripsi rincian kendala wajib diisi")
	}

	// Normalisasi prioritas
	switch req.Priority {
	case domain.PriorityP1, domain.PriorityP2, domain.PriorityP3:
		// Valid
	default:
		req.Priority = domain.PriorityP3
	}

	ticket := &domain.Ticket{
		ClientID:       currentUser.ID,
		Title:          strings.TrimSpace(req.Title),
		Description:    strings.TrimSpace(req.Description),
		Priority:       req.Priority,
		IsMinorFeature: req.IsMinorFeature,
	}

	return u.ticketRepo.CreateTicket(ctx, ticket)
}

func (u *ticketUsecase) UpdateStatus(ctx context.Context, currentUser *domain.User, req domain.UpdateTicketStatusRequest) error {
	if u.ticketRepo == nil {
		return errors.New("koneksi basis data belum siap")
	}

	// Hanya Engineer, Direktur, dan Admin yang boleh merubah status tiket SLA
	if currentUser.Role != domain.RoleEngineer && currentUser.Role != domain.RoleDirektur && currentUser.Role != domain.RoleAdmin {
		return errors.New("hanya teknisi engineer atau direksi yang dapat mengubah status tiket")
	}

	if req.ID <= 0 {
		return errors.New("ID tiket tidak valid")
	}

	return u.ticketRepo.UpdateStatus(ctx, req.ID, req.Status)
}
