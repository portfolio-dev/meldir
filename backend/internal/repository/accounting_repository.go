package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"meldir-backend/internal/domain"
)

type AccountingRepository interface {
	ListJournals(ctx context.Context, limit int) ([]domain.JournalEntry, error)
	CreateJournal(ctx context.Context, entry *domain.JournalEntry) (*domain.JournalEntry, error)
	GetNextJournalNumber(ctx context.Context) (string, error)
}

type accountingRepo struct {
	pool *pgxpool.Pool
}

func NewAccountingRepository(pool *pgxpool.Pool) AccountingRepository {
	return &accountingRepo{pool: pool}
}

func (r *accountingRepo) GetNextJournalNumber(ctx context.Context) (string, error) {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	var count int64
	err := r.pool.QueryRow(ctx, "SELECT COUNT(*) FROM accounting_journals").Scan(&count)
	if err != nil {
		return fmt.Sprintf("JU-%d/%02d/%03d", year, month, 1), nil
	}
	return fmt.Sprintf("JU-%d/%02d/%03d", year, month, count+1), nil
}

func (r *accountingRepo) ListJournals(ctx context.Context, limit int) ([]domain.JournalEntry, error) {
	if limit <= 0 {
		limit = 100
	}

	query := `
		SELECT j.id,
		       COALESCE(j.journal_number, ''),
		       COALESCE(j.journal_date::text, ''),
		       COALESCE(j.source_type::text, 'manual_adjustment'),
		       COALESCE(j.source_reference_id, ''),
		       COALESCE(j.memo, ''),
		       COALESCE((
		           SELECT a.account_name 
		           FROM accounting_journal_items ji
		           JOIN accounting_chart_of_accounts a ON ji.account_id = a.id
		           WHERE ji.journal_id = j.id AND ji.debit_amount > 0
		           ORDER BY ji.id ASC LIMIT 1
		       ), '11010 - Kas Operasional / Petty Cash') AS debit_account,
		       COALESCE((
		           SELECT a.account_name 
		           FROM accounting_journal_items ji
		           JOIN accounting_chart_of_accounts a ON ji.account_id = a.id
		           WHERE ji.journal_id = j.id AND ji.credit_amount > 0
		           ORDER BY ji.id ASC LIMIT 1
		       ), '41010 - Pendapatan Jasa Managed Maintenance Care') AS credit_account,
		       COALESCE(j.total_debit, 0.0)::float8,
		       COALESCE(j.total_credit, 0.0)::float8,
		       COALESCE(j.is_posted, true),
		       COALESCE(j.created_by, 0),
		       COALESCE(j.created_at, CURRENT_TIMESTAMP)
		FROM accounting_journals j
		ORDER BY j.journal_date DESC, j.created_at DESC
		LIMIT $1
	`

	rows, err := r.pool.Query(ctx, query, limit)
	if err != nil {
		log.Printf("⚠️ AccountingRepo.ListJournals main query error: %v, falling back to direct query", err)
		fallbackQuery := `
			SELECT id,
			       COALESCE(journal_number, ''),
			       COALESCE(journal_date::text, ''),
			       COALESCE(source_type::text, 'manual_adjustment'),
			       COALESCE(source_reference_id, ''),
			       COALESCE(memo, ''),
			       '11010 - Kas Operasional / Petty Cash' AS debit_account,
			       '41010 - Pendapatan Jasa Managed Maintenance Care' AS credit_account,
			       COALESCE(total_debit, 0.0)::float8,
			       COALESCE(total_credit, 0.0)::float8,
			       COALESCE(is_posted, true),
			       COALESCE(created_by, 0),
			       COALESCE(created_at, CURRENT_TIMESTAMP)
			FROM accounting_journals
			ORDER BY journal_date DESC, created_at DESC
			LIMIT $1
		`
		rows, err = r.pool.Query(ctx, fallbackQuery, limit)
		if err != nil {
			log.Printf("❌ AccountingRepo.ListJournals fallback query error: %v", err)
			return nil, err
		}
	}
	defer rows.Close()

	entries := make([]domain.JournalEntry, 0)
	for rows.Next() {
		var j domain.JournalEntry
		err := rows.Scan(
			&j.ID, &j.JournalNumber, &j.JournalDate, &j.SourceType,
			&j.SourceReferenceID, &j.Memo,
			&j.DebitAccount, &j.CreditAccount,
			&j.TotalDebit, &j.TotalCredit, &j.IsPosted,
			&j.CreatedBy, &j.CreatedAt,
		)
		if err != nil {
			log.Printf("❌ AccountingRepo.ListJournals scan error: %v", err)
			return nil, err
		}
		j.Amount = j.TotalDebit
		entries = append(entries, j)
	}

	return entries, nil
}

func (r *accountingRepo) CreateJournal(ctx context.Context, entry *domain.JournalEntry) (*domain.JournalEntry, error) {
	if entry.JournalNumber == "" {
		num, err := r.GetNextJournalNumber(ctx)
		if err != nil {
			return nil, err
		}
		entry.JournalNumber = num
	}

	entry.TotalDebit = entry.Amount
	entry.TotalCredit = entry.Amount
	entry.IsPosted = true
	entry.CreatedAt = time.Now().UTC()

	switch entry.SourceType {
	case "invoice_payment", "engineer_payout", "corporate_expense", "tax_settlement", "opening_balance":
		// valid enum values
	default:
		entry.SourceType = "manual_adjustment"
	}

	var createdByVal int64 = entry.CreatedBy
	if createdByVal <= 0 {
		_ = r.pool.QueryRow(ctx, "SELECT id FROM users ORDER BY id ASC LIMIT 1").Scan(&createdByVal)
		if createdByVal <= 0 {
			createdByVal = 1
		}
	}
	entry.CreatedBy = createdByVal

	query := `
		INSERT INTO accounting_journals (
			journal_number, journal_date, source_type, source_reference_id,
			memo, total_debit, total_credit, is_posted, created_by, created_at
		) VALUES (
			$1, $2::date, $3, $4, $5, $6, $7, $8, $9, $10
		) RETURNING id
	`

	err := r.pool.QueryRow(
		ctx, query,
		entry.JournalNumber, entry.JournalDate, entry.SourceType, entry.SourceReferenceID,
		entry.Memo, entry.TotalDebit, entry.TotalCredit, entry.IsPosted, entry.CreatedBy, entry.CreatedAt,
	).Scan(&entry.ID)
	if err != nil {
		log.Printf("❌ AccountingRepo.CreateJournal insert error: %v", err)
		return nil, err
	}

	var debitAccID int64
	var creditAccID int64

	if entry.DebitAccount != "" {
		_ = r.pool.QueryRow(ctx, "SELECT id FROM accounting_chart_of_accounts WHERE account_name ILIKE $1 OR account_code = $2 ORDER BY id ASC LIMIT 1", "%"+entry.DebitAccount+"%", entry.DebitAccount).Scan(&debitAccID)
	}
	if debitAccID == 0 {
		_ = r.pool.QueryRow(ctx, "SELECT id FROM accounting_chart_of_accounts WHERE category = 'asset' ORDER BY id ASC LIMIT 1").Scan(&debitAccID)
	}

	if entry.CreditAccount != "" {
		_ = r.pool.QueryRow(ctx, "SELECT id FROM accounting_chart_of_accounts WHERE account_name ILIKE $1 OR account_code = $2 ORDER BY id ASC LIMIT 1", "%"+entry.CreditAccount+"%", entry.CreditAccount).Scan(&creditAccID)
	}
	if creditAccID == 0 {
		_ = r.pool.QueryRow(ctx, "SELECT id FROM accounting_chart_of_accounts WHERE category = 'revenue' ORDER BY id ASC LIMIT 1").Scan(&creditAccID)
	}

	if debitAccID > 0 {
		_, _ = r.pool.Exec(ctx, "INSERT INTO accounting_journal_items (journal_id, account_id, debit_amount, credit_amount, description, created_at) VALUES ($1, $2, $3, 0.00, $4, $5)", entry.ID, debitAccID, entry.TotalDebit, entry.Memo, entry.CreatedAt)
	}
	if creditAccID > 0 {
		_, _ = r.pool.Exec(ctx, "INSERT INTO accounting_journal_items (journal_id, account_id, debit_amount, credit_amount, description, created_at) VALUES ($1, $2, 0.00, $3, $4, $5)", entry.ID, creditAccID, entry.TotalCredit, entry.Memo, entry.CreatedAt)
	}

	return entry, nil
}
