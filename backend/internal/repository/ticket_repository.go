package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"meldir-backend/internal/domain"
)

type TicketRepository interface {
	ListTickets(ctx context.Context, filter domain.TicketFilter) ([]domain.Ticket, error)
	FindByID(ctx context.Context, id int64) (*domain.Ticket, error)
	CreateTicket(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error)
	UpdateStatus(ctx context.Context, id int64, status domain.TicketStatus) error
	GetNextTicketCode(ctx context.Context) (string, error)
}

type ticketRepo struct {
	pool *pgxpool.Pool
}

func NewTicketRepository(pool *pgxpool.Pool) TicketRepository {
	return &ticketRepo{pool: pool}
}

func (r *ticketRepo) GetNextTicketCode(ctx context.Context) (string, error) {
	year := time.Now().Year()
	var count int64
	err := r.pool.QueryRow(ctx, "SELECT COUNT(*) FROM tickets").Scan(&count)
	if err != nil {
		return fmt.Sprintf("TKT-%d-%03d", year, 1), nil
	}
	return fmt.Sprintf("TKT-%d-%03d", year, count+1), nil
}

func (r *ticketRepo) ListTickets(ctx context.Context, filter domain.TicketFilter) ([]domain.Ticket, error) {
	query := `
		SELECT t.id, t.ticket_code, t.project_id, COALESCE(p.project_name, ''),
		       t.client_id, u.name, u.email, t.contract_id,
		       t.title, t.description, t.priority::text, t.is_minor_feature,
		       t.status::text, t.assigned_engineer_id, COALESCE(eng.name, ''),
		       t.sla_deadline, t.is_sla_breached,
		       COALESCE((SELECT SUM(hours_spent) FROM timesheet_logs WHERE ticket_id = t.id), 0.0),
		       t.resolved_at, t.created_at, t.updated_at
		FROM tickets t
		JOIN users u ON t.client_id = u.id
		LEFT JOIN projects p ON t.project_id = p.id
		LEFT JOIN users eng ON t.assigned_engineer_id = eng.id
		WHERE ($1 = 0 OR t.client_id = $1)
		  AND ($2 = 0 OR t.assigned_engineer_id = $2)
		  AND ($3 = '' OR t.status::text = $3)
		  AND ($4 = '' OR t.priority::text = $4)
		  AND ($5 = '' OR LOWER(t.title) LIKE '%' || LOWER($5) || '%' OR LOWER(t.ticket_code) LIKE '%' || LOWER($5) || '%')
		ORDER BY t.created_at DESC
	`

	search := strings.TrimSpace(filter.Search)
	rows, err := r.pool.Query(ctx, query, filter.ClientID, filter.EngineerID, filter.Status, filter.Priority, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := make([]domain.Ticket, 0)
	for rows.Next() {
		var t domain.Ticket
		var priorityStr, statusStr string
		err := rows.Scan(
			&t.ID, &t.TicketCode, &t.ProjectID, &t.ProjectName,
			&t.ClientID, &t.ClientName, &t.ClientEmail, &t.ContractID,
			&t.Title, &t.Description, &priorityStr, &t.IsMinorFeature,
			&statusStr, &t.AssignedEngineerID, &t.EngineerName,
			&t.SLADeadline, &t.IsSLABreached, &t.HoursSpent,
			&t.ResolvedAt, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t.Priority = domain.TicketPriority(priorityStr)
		t.Status = domain.TicketStatus(statusStr)
		tickets = append(tickets, t)
	}

	return tickets, nil
}

func (r *ticketRepo) FindByID(ctx context.Context, id int64) (*domain.Ticket, error) {
	query := `
		SELECT t.id, t.ticket_code, t.project_id, COALESCE(p.project_name, ''),
		       t.client_id, u.name, u.email, t.contract_id,
		       t.title, t.description, t.priority::text, t.is_minor_feature,
		       t.status::text, t.assigned_engineer_id, COALESCE(eng.name, ''),
		       t.sla_deadline, t.is_sla_breached,
		       COALESCE((SELECT SUM(hours_spent) FROM timesheet_logs WHERE ticket_id = t.id), 0.0),
		       t.resolved_at, t.created_at, t.updated_at
		FROM tickets t
		JOIN users u ON t.client_id = u.id
		LEFT JOIN projects p ON t.project_id = p.id
		LEFT JOIN users eng ON t.assigned_engineer_id = eng.id
		WHERE t.id = $1
		LIMIT 1
	`
	var t domain.Ticket
	var priorityStr, statusStr string
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.TicketCode, &t.ProjectID, &t.ProjectName,
		&t.ClientID, &t.ClientName, &t.ClientEmail, &t.ContractID,
		&t.Title, &t.Description, &priorityStr, &t.IsMinorFeature,
		&statusStr, &t.AssignedEngineerID, &t.EngineerName,
		&t.SLADeadline, &t.IsSLABreached, &t.HoursSpent,
		&t.ResolvedAt, &t.CreatedAt, &t.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("tiket tidak ditemukan")
		}
		return nil, err
	}
	t.Priority = domain.TicketPriority(priorityStr)
	t.Status = domain.TicketStatus(statusStr)
	return &t, nil
}

func (r *ticketRepo) CreateTicket(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	code, err := r.GetNextTicketCode(ctx)
	if err != nil {
		return nil, err
	}
	ticket.TicketCode = code

	// Hitung batas SLA berdasarkan prioritas
	now := time.Now().UTC()
	var deadline time.Time
	switch ticket.Priority {
	case domain.PriorityP1:
		deadline = now.Add(4 * time.Hour)
	case domain.PriorityP2:
		deadline = now.Add(12 * time.Hour)
	default:
		deadline = now.Add(48 * time.Hour)
	}
	ticket.SLADeadline = deadline
	ticket.Status = domain.TicketStatusOpen
	ticket.CreatedAt = now
	ticket.UpdatedAt = now

	query := `
		INSERT INTO tickets (
			ticket_code, client_id, title, description, priority,
			is_minor_feature, status, sla_deadline, is_sla_breached,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, false, $9, $10
		) RETURNING id
	`
	err = r.pool.QueryRow(
		ctx, query,
		ticket.TicketCode, ticket.ClientID, ticket.Title, ticket.Description,
		string(ticket.Priority), ticket.IsMinorFeature, string(ticket.Status),
		ticket.SLADeadline, ticket.CreatedAt, ticket.UpdatedAt,
	).Scan(&ticket.ID)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *ticketRepo) UpdateStatus(ctx context.Context, id int64, status domain.TicketStatus) error {
	now := time.Now().UTC()
	var query string
	var err error

	if status == domain.TicketStatusResolved || status == domain.TicketStatusClosed {
		query = "UPDATE tickets SET status = $1, resolved_at = $2, updated_at = $3 WHERE id = $4"
		_, err = r.pool.Exec(ctx, query, string(status), now, now, id)
	} else {
		query = "UPDATE tickets SET status = $1, resolved_at = NULL, updated_at = $2 WHERE id = $3"
		_, err = r.pool.Exec(ctx, query, string(status), now, id)
	}
	return err
}
