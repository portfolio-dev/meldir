package usecase

import (
	"context"
	"errors"
	"strings"

	"meldir-backend/internal/domain"
	"meldir-backend/internal/repository"
)

type LeadUsecase interface {
	ListLeads(ctx context.Context, status string) ([]domain.Lead, error)
	CreatePublicLead(ctx context.Context, req domain.CreateLeadRequest) (*domain.Lead, error)
	UpdateStatus(ctx context.Context, req domain.UpdateLeadStatusRequest) error
}

type leadUsecase struct {
	leadRepo repository.LeadRepository
}

func NewLeadUsecase(leadRepo repository.LeadRepository) LeadUsecase {
	return &leadUsecase{leadRepo: leadRepo}
}

func (u *leadUsecase) ListLeads(ctx context.Context, status string) ([]domain.Lead, error) {
	if u.leadRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}
	return u.leadRepo.ListLeads(ctx, strings.TrimSpace(status))
}

func (u *leadUsecase) CreatePublicLead(ctx context.Context, req domain.CreateLeadRequest) (*domain.Lead, error) {
	if u.leadRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("nama lengkap atau penanggung jawab wajib diisi")
	}

	phone := strings.TrimSpace(req.PhoneWA)
	if phone == "" {
		return nil, errors.New("nomor kontak WhatsApp aktif wajib diisi")
	}

	lead := &domain.Lead{
		Name:            name,
		Email:           strings.TrimSpace(req.Email),
		PhoneWA:         phone,
		CompanyName:     strings.TrimSpace(req.CompanyName),
		ServiceInterest: strings.TrimSpace(req.ServiceInterest),
		BudgetRange:     strings.TrimSpace(req.BudgetRange),
		Message:         strings.TrimSpace(req.Message),
		Source:          strings.TrimSpace(req.Source),
	}

	return u.leadRepo.CreateLead(ctx, lead)
}

func (u *leadUsecase) UpdateStatus(ctx context.Context, req domain.UpdateLeadStatusRequest) error {
	if u.leadRepo == nil {
		return errors.New("koneksi basis data belum siap")
	}
	if req.ID <= 0 {
		return errors.New("ID lead tidak valid")
	}
	if req.Status == "" {
		return errors.New("status prospek wajib diisi")
	}

	return u.leadRepo.UpdateLeadStatus(ctx, req.ID, req.Status, strings.TrimSpace(req.Notes))
}
