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
		SELECT COALESCE(id, 0),
		       COALESCE(journal_number, ''),
		       COALESCE(journal_date::text, ''),
		       COALESCE(source_type::text, 'general_entry'),
		       COALESCE(source_reference_id, ''),
		       COALESCE(memo, ''),
		       COALESCE(debit_account, ''),
		       COALESCE(credit_account, ''),
		       COALESCE(total_debit, 0.0)::float8,
		       COALESCE(total_credit, 0.0)::float8,
		       COALESCE(is_posted, true),
		       COALESCE(created_by, 0),
		       COALESCE(created_at, CURRENT_TIMESTAMP)
		FROM accounting_journals
		ORDER BY journal_date DESC, created_at DESC
		LIMIT $1
	`

	rows, err := r.pool.Query(ctx, query, limit)
	if err != nil {
		log.Printf("❌ AccountingRepo.ListJournals query error: %v", err)
		return nil, err
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
	if entry.SourceType == "" {
		entry.SourceType = "general_entry"
	}

	query := `
		INSERT INTO accounting_journals (
			journal_number, journal_date, source_type, source_reference_id,
			memo, debit_account, credit_account, total_debit, total_credit,
			is_posted, created_by, created_at
		) VALUES (
			$1, $2::date, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		) RETURNING id
	`

	var createdByVal interface{} = entry.CreatedBy
	if entry.CreatedBy <= 0 {
		createdByVal = nil
	}

	err := r.pool.QueryRow(
		ctx, query,
		entry.JournalNumber, entry.JournalDate, entry.SourceType, entry.SourceReferenceID,
		entry.Memo, entry.DebitAccount, entry.CreditAccount,
		entry.TotalDebit, entry.TotalCredit, entry.IsPosted, createdByVal, entry.CreatedAt,
	).Scan(&entry.ID)
	if err != nil {
		log.Printf("❌ AccountingRepo.CreateJournal insert error: %v", err)
		return nil, err
	}

	return entry, nil
}
