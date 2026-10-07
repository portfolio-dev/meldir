package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"meldir-backend/internal/domain"
	"meldir-backend/internal/repository"
)

type AccountingUsecase interface {
	ListJournals(ctx context.Context, currentUser *domain.User, limit int) ([]domain.JournalEntry, error)
	CreateJournal(ctx context.Context, currentUser *domain.User, req domain.CreateJournalRequest) (*domain.JournalEntry, error)
}

type accountingUsecase struct {
	accountingRepo repository.AccountingRepository
}

func NewAccountingUsecase(accountingRepo repository.AccountingRepository) AccountingUsecase {
	return &accountingUsecase{accountingRepo: accountingRepo}
}

func (u *accountingUsecase) ListJournals(ctx context.Context, currentUser *domain.User, limit int) ([]domain.JournalEntry, error) {
	if u.accountingRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	if currentUser.Role != domain.RoleDirektur && currentUser.Role != domain.RoleAdmin && currentUser.Role != domain.RoleAudit {
		return nil, errors.New("Anda tidak memiliki hak akses ke Buku Besar akuntansi")
	}

	return u.accountingRepo.ListJournals(ctx, limit)
}

func (u *accountingUsecase) CreateJournal(ctx context.Context, currentUser *domain.User, req domain.CreateJournalRequest) (*domain.JournalEntry, error) {
	if u.accountingRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	if currentUser.Role != domain.RoleDirektur && currentUser.Role != domain.RoleAdmin {
		return nil, errors.New("hanya direktur atau admin akuntansi yang dapat memposting jurnal umum")
	}

	if req.Amount <= 0 {
		return nil, errors.New("nominal transaksi jurnal harus lebih dari 0")
	}

	if strings.TrimSpace(req.Description) == "" {
		return nil, errors.New("keterangan transaksi wajib diisi")
	}

	if strings.TrimSpace(req.DebitAccount) == "" || strings.TrimSpace(req.CreditAccount) == "" {
		return nil, errors.New("posisi akun debit dan kredit berpasangan wajib ditentukan")
	}

	if strings.TrimSpace(req.Date) == "" {
		req.Date = time.Now().Format("2006-01-02")
	}

	entry := &domain.JournalEntry{
		JournalDate:   req.Date,
		Memo:          strings.TrimSpace(req.Description),
		DebitAccount:  strings.TrimSpace(req.DebitAccount),
		CreditAccount: strings.TrimSpace(req.CreditAccount),
		Amount:        req.Amount,
		CreatedBy:     currentUser.ID,
	}

	return u.accountingRepo.CreateJournal(ctx, entry)
}
