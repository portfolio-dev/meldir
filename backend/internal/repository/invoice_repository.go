package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"meldir-backend/internal/domain"
)

type InvoiceRepository interface {
	ListInvoices(ctx context.Context, clientID int64) ([]domain.Invoice, error)
	FindByID(ctx context.Context, id int64) (*domain.Invoice, error)
	CreateInvoice(ctx context.Context, inv *domain.Invoice) (*domain.Invoice, error)
	UpdateStatus(ctx context.Context, id int64, status domain.InvoiceStatus) error
	GetNextInvoiceNumber(ctx context.Context) (string, error)
}

type invoiceRepo struct {
	pool *pgxpool.Pool
}

func NewInvoiceRepository(pool *pgxpool.Pool) InvoiceRepository {
	return &invoiceRepo{pool: pool}
}

func (r *invoiceRepo) GetNextInvoiceNumber(ctx context.Context) (string, error) {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	var count int64
	err := r.pool.QueryRow(ctx, "SELECT COUNT(*) FROM invoices").Scan(&count)
	if err != nil {
		return fmt.Sprintf("INV/%d/%02d/%03d", year, month, 1), nil
	}
	return fmt.Sprintf("INV/%d/%02d/%03d", year, month, count+1), nil
}

func (r *invoiceRepo) ListInvoices(ctx context.Context, clientID int64) ([]domain.Invoice, error) {
	query := `
		SELECT COALESCE(i.id, 0),
		       COALESCE(i.invoice_number, ''),
		       COALESCE(i.client_id, 0),
		       COALESCE(u.name, COALESCE(i.client_name, ''), ''),
		       i.contract_id,
		       COALESCE(i.amount, 0.0)::float8,
		       COALESCE(i.tax_amount, 0.0)::float8,
		       (COALESCE(i.amount, 0.0) + COALESCE(i.tax_amount, 0.0))::float8,
		       COALESCE(i.due_date::text, ''),
		       COALESCE(i.status::text, 'unpaid'),
		       COALESCE(i.bank_destination, ''),
		       i.tax_invoice_number,
		       i.paid_at,
		       COALESCE(i.created_at, CURRENT_TIMESTAMP)
		FROM invoices i
		LEFT JOIN users u ON i.client_id = u.id
		WHERE ($1 = 0 OR i.client_id = $1)
		ORDER BY i.created_at DESC
	`

	rows, err := r.pool.Query(ctx, query, clientID)
	if err != nil {
		log.Printf("❌ InvoiceRepo.ListInvoices query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	invoices := make([]domain.Invoice, 0)
	for rows.Next() {
		var inv domain.Invoice
		var statusStr string
		err := rows.Scan(
			&inv.ID, &inv.InvoiceNumber, &inv.ClientID, &inv.ClientName,
			&inv.ContractID, &inv.Amount, &inv.TaxAmount, &inv.TotalAmount,
			&inv.DueDate, &statusStr, &inv.BankDestination, &inv.TaxInvoiceNumber,
			&inv.PaidAt, &inv.CreatedAt,
		)
		if err != nil {
			log.Printf("❌ InvoiceRepo.ListInvoices scan error: %v", err)
			return nil, err
		}
		inv.Status = domain.InvoiceStatus(statusStr)
		invoices = append(invoices, inv)
	}

	return invoices, nil
}

func (r *invoiceRepo) FindByID(ctx context.Context, id int64) (*domain.Invoice, error) {
	query := `
		SELECT COALESCE(i.id, 0),
		       COALESCE(i.invoice_number, ''),
		       COALESCE(i.client_id, 0),
		       COALESCE(u.name, COALESCE(i.client_name, ''), ''),
		       i.contract_id,
		       COALESCE(i.amount, 0.0)::float8,
		       COALESCE(i.tax_amount, 0.0)::float8,
		       (COALESCE(i.amount, 0.0) + COALESCE(i.tax_amount, 0.0))::float8,
		       COALESCE(i.due_date::text, ''),
		       COALESCE(i.status::text, 'unpaid'),
		       COALESCE(i.bank_destination, ''),
		       i.tax_invoice_number,
		       i.paid_at,
		       COALESCE(i.created_at, CURRENT_TIMESTAMP)
		FROM invoices i
		LEFT JOIN users u ON i.client_id = u.id
		WHERE i.id = $1
		LIMIT 1
	`
	var inv domain.Invoice
	var statusStr string
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&inv.ID, &inv.InvoiceNumber, &inv.ClientID, &inv.ClientName,
		&inv.ContractID, &inv.Amount, &inv.TaxAmount, &inv.TotalAmount,
		&inv.DueDate, &statusStr, &inv.BankDestination, &inv.TaxInvoiceNumber,
		&inv.PaidAt, &inv.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("invoice tidak ditemukan")
		}
		log.Printf("❌ InvoiceRepo.FindByID scan error: %v", err)
		return nil, err
	}
	inv.Status = domain.InvoiceStatus(statusStr)
	return &inv, nil
}

func (r *invoiceRepo) CreateInvoice(ctx context.Context, inv *domain.Invoice) (*domain.Invoice, error) {
	num, err := r.GetNextInvoiceNumber(ctx)
	if err != nil {
		return nil, err
	}
	inv.InvoiceNumber = num
	inv.TaxAmount = float64(int64(inv.Amount*0.11 + 0.5)) // PPN 11% pembulatan
	inv.TotalAmount = inv.Amount + inv.TaxAmount
	inv.Status = domain.InvoiceStatusUnpaid
	inv.CreatedAt = time.Now().UTC()
	if inv.BankDestination == "" {
		inv.BankDestination = "PT. Melayani Digital Raya - Bank Mandiri & BCA"
	}

	query := `
		INSERT INTO invoices (
			invoice_number, client_id, client_name, amount, tax_amount,
			due_date, status, bank_destination, created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6::date, $7, $8, $9
		) RETURNING id
	`

	var clientIDVal interface{} = inv.ClientID
	if inv.ClientID <= 0 {
		clientIDVal = nil
	}

	err = r.pool.QueryRow(
		ctx, query,
		inv.InvoiceNumber, clientIDVal, inv.ClientName, inv.Amount, inv.TaxAmount,
		inv.DueDate, string(inv.Status), inv.BankDestination, inv.CreatedAt,
	).Scan(&inv.ID)
	if err != nil {
		log.Printf("❌ InvoiceRepo.CreateInvoice insert error: %v", err)
		return nil, err
	}

	return inv, nil
}

func (r *invoiceRepo) UpdateStatus(ctx context.Context, id int64, status domain.InvoiceStatus) error {
	now := time.Now().UTC()
	var err error
	if status == domain.InvoiceStatusPaid {
		_, err = r.pool.Exec(ctx, "UPDATE invoices SET status = $1, paid_at = $2 WHERE id = $3", string(status), now, id)
	} else {
		_, err = r.pool.Exec(ctx, "UPDATE invoices SET status = $1, paid_at = NULL WHERE id = $2", string(status), id)
	}
	if err != nil {
		log.Printf("❌ InvoiceRepo.UpdateStatus error: %v", err)
	}
	return err
}
