package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"meldir-backend/internal/domain"
	"meldir-backend/internal/repository"
)

type InvoiceUsecase interface {
	ListInvoices(ctx context.Context, currentUser *domain.User) ([]domain.Invoice, error)
	GetInvoiceByID(ctx context.Context, currentUser *domain.User, id int64) (*domain.Invoice, error)
	CreateInvoice(ctx context.Context, currentUser *domain.User, req domain.CreateInvoiceRequest) (*domain.Invoice, error)
	UpdateStatus(ctx context.Context, currentUser *domain.User, req domain.UpdateInvoiceStatusRequest) error
}

type invoiceUsecase struct {
	invoiceRepo    repository.InvoiceRepository
	accountingRepo repository.AccountingRepository
}

func NewInvoiceUsecase(
	invoiceRepo repository.InvoiceRepository,
	accountingRepo repository.AccountingRepository,
) InvoiceUsecase {
	return &invoiceUsecase{
		invoiceRepo:    invoiceRepo,
		accountingRepo: accountingRepo,
	}
}

func (u *invoiceUsecase) ListInvoices(ctx context.Context, currentUser *domain.User) ([]domain.Invoice, error) {
	if u.invoiceRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	var clientID int64
	if currentUser.Role == domain.RoleKlien {
		clientID = currentUser.ID
	}

	return u.invoiceRepo.ListInvoices(ctx, clientID)
}

func (u *invoiceUsecase) GetInvoiceByID(ctx context.Context, currentUser *domain.User, id int64) (*domain.Invoice, error) {
	if u.invoiceRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	inv, err := u.invoiceRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if currentUser.Role == domain.RoleKlien && inv.ClientID != currentUser.ID {
		return nil, errors.New("Anda tidak memiliki akses ke faktur ini")
	}

	return inv, nil
}

func (u *invoiceUsecase) CreateInvoice(ctx context.Context, currentUser *domain.User, req domain.CreateInvoiceRequest) (*domain.Invoice, error) {
	if u.invoiceRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	if currentUser.Role != domain.RoleDirektur && currentUser.Role != domain.RoleAdmin {
		return nil, errors.New("hanya direktur atau admin yang berhak menerbitkan faktur invoice")
	}

	if req.Amount <= 0 {
		return nil, errors.New("nominal DPP pokok faktur harus lebih dari 0")
	}

	if strings.TrimSpace(req.ClientName) == "" {
		return nil, errors.New("nama perusahaan / klien penerima tagihan wajib diisi")
	}

	if req.ClientID <= 0 {
		req.ClientID = currentUser.ID
	}

	if strings.TrimSpace(req.DueDate) == "" {
		req.DueDate = time.Now().Add(14 * 24 * time.Hour).Format("2006-01-02")
	}

	inv := &domain.Invoice{
		ClientID:   req.ClientID,
		ClientName: strings.TrimSpace(req.ClientName),
		Amount:     req.Amount,
		DueDate:    req.DueDate,
	}

	createdInv, err := u.invoiceRepo.CreateInvoice(ctx, inv)
	if err != nil {
		return nil, err
	}

	// Otomasi pembukuan double-entry SAK EMKM: Piutang (Debit) vs Pendapatan & PPN (Kredit)
	if u.accountingRepo != nil {
		_ = u.accountingRepo.CreateJournal(ctx, &domain.JournalEntry{
			JournalDate:       time.Now().Format("2006-01-02"),
			SourceType:        "invoice_issued",
			SourceReferenceID: createdInv.InvoiceNumber,
			Memo:              "Penerbitan tagihan " + createdInv.InvoiceNumber + " a/n " + createdInv.ClientName,
			DebitAccount:      "1-1201 Piutang Usaha Klien",
			CreditAccount:     "4-1002 Pendapatan Kontrak Managed Care SLA",
			Amount:            createdInv.TotalAmount,
			CreatedBy:         currentUser.ID,
		})
	}

	return createdInv, nil
}

func (u *invoiceUsecase) UpdateStatus(ctx context.Context, currentUser *domain.User, req domain.UpdateInvoiceStatusRequest) error {
	if u.invoiceRepo == nil {
		return errors.New("koneksi basis data belum siap")
	}

	if currentUser.Role != domain.RoleDirektur && currentUser.Role != domain.RoleAdmin {
		return errors.New("hanya direktur atau admin yang dapat merubah status pembayaran faktur")
	}

	inv, err := u.invoiceRepo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	err = u.invoiceRepo.UpdateStatus(ctx, req.ID, req.Status)
	if err != nil {
		return err
	}

	// Jika status diubah menjadi PAID, otomatis bukukan penerimaan Kas/Bank (Debit) vs Piutang (Kredit)
	if req.Status == domain.InvoiceStatusPaid && u.accountingRepo != nil {
		_ = u.accountingRepo.CreateJournal(ctx, &domain.JournalEntry{
			JournalDate:       time.Now().Format("2006-01-02"),
			SourceType:        "invoice_payment",
			SourceReferenceID: inv.InvoiceNumber,
			Memo:              "Penerimaan pelunasan tagihan " + inv.InvoiceNumber + " a/n " + inv.ClientName,
			DebitAccount:      "1-1002 Bank BCA Utama Korporat",
			CreditAccount:     "1-1201 Piutang Usaha Klien",
			Amount:            inv.TotalAmount,
			CreatedBy:         currentUser.ID,
		})
	}

	return nil
}
