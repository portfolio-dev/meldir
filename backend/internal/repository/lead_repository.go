package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"meldir-backend/internal/domain"
)

type LeadRepository interface {
	ListLeads(ctx context.Context, status string) ([]domain.Lead, error)
	CreateLead(ctx context.Context, lead *domain.Lead) (*domain.Lead, error)
	UpdateLeadStatus(ctx context.Context, id int64, status domain.LeadStatus, notes string) error
	GetNextLeadCode(ctx context.Context) (string, error)
}

type leadRepo struct {
	pool *pgxpool.Pool
}

func NewLeadRepository(pool *pgxpool.Pool) LeadRepository {
	return &leadRepo{pool: pool}
}

func (r *leadRepo) GetNextLeadCode(ctx context.Context) (string, error) {
	now := time.Now()
	year := now.Year()
	var count int64
	err := r.pool.QueryRow(ctx, "SELECT COUNT(*) FROM inbound_leads").Scan(&count)
	if err != nil {
		return fmt.Sprintf("LEAD-%d-%03d", year, 1), nil
	}
	return fmt.Sprintf("LEAD-%d-%03d", year, count+1), nil
}

func (r *leadRepo) ListLeads(ctx context.Context, status string) ([]domain.Lead, error) {
	query := `
		SELECT COALESCE(id, 0),
		       COALESCE(lead_code, ''),
		       COALESCE(name, ''),
		       COALESCE(email, ''),
		       COALESCE(phone_wa, ''),
		       COALESCE(company_name, ''),
		       COALESCE(service_interest::text, 'managed_care'),
		       COALESCE(budget_range, ''),
		       COALESCE(message, ''),
		       COALESCE(status::text, 'new'),
		       COALESCE(source, 'website_landing'),
		       assigned_admin_id,
		       COALESCE(notes, ''),
		       COALESCE(created_at, CURRENT_TIMESTAMP),
		       COALESCE(updated_at, CURRENT_TIMESTAMP)
		FROM inbound_leads
		WHERE ($1 = '' OR status::text = $1)
		ORDER BY created_at DESC
	`

	rows, err := r.pool.Query(ctx, query, status)
	if err != nil {
		log.Printf("❌ LeadRepo.ListLeads query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	leads := make([]domain.Lead, 0)
	for rows.Next() {
		var l domain.Lead
		var statusStr string
		err := rows.Scan(
			&l.ID, &l.LeadCode, &l.Name, &l.Email, &l.PhoneWA,
			&l.CompanyName, &l.ServiceInterest, &l.BudgetRange, &l.Message,
			&statusStr, &l.Source, &l.AssignedAdminID, &l.Notes,
			&l.CreatedAt, &l.UpdatedAt,
		)
		if err != nil {
			log.Printf("❌ LeadRepo.ListLeads scan error: %v", err)
			return nil, err
		}
		l.Status = domain.LeadStatus(statusStr)
		leads = append(leads, l)
	}

	return leads, nil
}

func (r *leadRepo) CreateLead(ctx context.Context, lead *domain.Lead) (*domain.Lead, error) {
	if lead.LeadCode == "" {
		code, err := r.GetNextLeadCode(ctx)
		if err != nil {
			return nil, err
		}
		lead.LeadCode = code
	}

	lead.Status = domain.LeadStatusNew
	now := time.Now().UTC()
	lead.CreatedAt = now
	lead.UpdatedAt = now

	if lead.Source == "" {
		lead.Source = "website_landing"
	}
	if lead.ServiceInterest == "" {
		lead.ServiceInterest = "managed_care"
	}

	query := `
		INSERT INTO inbound_leads (
			lead_code, name, email, phone_wa, company_name,
			service_interest, budget_range, message, status, source,
			notes, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10,
			$11, $12, $13
		) RETURNING id
	`

	err := r.pool.QueryRow(
		ctx, query,
		lead.LeadCode, lead.Name, lead.Email, lead.PhoneWA, lead.CompanyName,
		lead.ServiceInterest, lead.BudgetRange, lead.Message, string(lead.Status), lead.Source,
		lead.Notes, lead.CreatedAt, lead.UpdatedAt,
	).Scan(&lead.ID)
	if err != nil {
		log.Printf("❌ LeadRepo.CreateLead insert error: %v", err)
		return nil, err
	}

	return lead, nil
}

func (r *leadRepo) UpdateLeadStatus(ctx context.Context, id int64, status domain.LeadStatus, notes string) error {
	now := time.Now().UTC()
	query := `
		UPDATE inbound_leads
		SET status = $1,
		    notes = CASE WHEN $2 != '' THEN $2 ELSE notes END,
		    updated_at = $3
		WHERE id = $4
	`
	_, err := r.pool.Exec(ctx, query, string(status), notes, now, id)
	if err != nil {
		log.Printf("❌ LeadRepo.UpdateLeadStatus error: %v", err)
	}
	return err
}
