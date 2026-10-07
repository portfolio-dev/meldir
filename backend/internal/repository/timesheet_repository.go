package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"meldir-backend/internal/domain"
)

type TimesheetRepository interface {
	ListTimesheets(ctx context.Context, engineerID int64, limit int) ([]domain.TimesheetLog, error)
	CreateTimesheet(ctx context.Context, log *domain.TimesheetLog) (*domain.TimesheetLog, error)
	GetSummary(ctx context.Context, engineerID int64) (*domain.TimesheetSummary, error)
}

type timesheetRepo struct {
	pool *pgxpool.Pool
}

func NewTimesheetRepository(pool *pgxpool.Pool) TimesheetRepository {
	return &timesheetRepo{pool: pool}
}

func (r *timesheetRepo) ListTimesheets(ctx context.Context, engineerID int64, limit int) ([]domain.TimesheetLog, error) {
	if limit <= 0 {
		limit = 50
	}

	query := `
		SELECT tl.id, tl.engineer_id, u.name, tl.project_id, COALESCE(tl.project_name, ''),
		       tl.ticket_id, COALESCE(tl.ticket_code, ''), tl.addon_order_id,
		       tl.hours_spent, tl.work_description, tl.log_date::text, tl.created_at
		FROM timesheet_logs tl
		JOIN users u ON tl.engineer_id = u.id
		WHERE ($1 = 0 OR tl.engineer_id = $1)
		ORDER BY tl.log_date DESC, tl.created_at DESC
		LIMIT $2
	`

	rows, err := r.pool.Query(ctx, query, engineerID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := make([]domain.TimesheetLog, 0)
	for rows.Next() {
		var l domain.TimesheetLog
		err := rows.Scan(
			&l.ID, &l.EngineerID, &l.EngineerName, &l.ProjectID, &l.ProjectName,
			&l.TicketID, &l.TicketCode, &l.AddonOrderID,
			&l.HoursSpent, &l.WorkDescription, &l.LogDate, &l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}

	return logs, nil
}

func (r *timesheetRepo) CreateTimesheet(ctx context.Context, log *domain.TimesheetLog) (*domain.TimesheetLog, error) {
	now := time.Now().UTC()
	log.CreatedAt = now

	// Jika ada ticket_code, cari ticket_id jika belum diset
	if log.TicketCode != "" && log.TicketID == nil {
		var tID int64
		err := r.pool.QueryRow(ctx, "SELECT id FROM tickets WHERE ticket_code = $1 LIMIT 1", log.TicketCode).Scan(&tID)
		if err == nil {
			log.TicketID = &tID
		}
	}

	query := `
		INSERT INTO timesheet_logs (
			engineer_id, ticket_id, ticket_code, project_name, hours_spent,
			work_description, log_date, created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7::date, $8
		) RETURNING id
	`

	err := r.pool.QueryRow(
		ctx, query,
		log.EngineerID, log.TicketID, log.TicketCode, log.ProjectName,
		log.HoursSpent, log.WorkDescription, log.LogDate, log.CreatedAt,
	).Scan(&log.ID)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (r *timesheetRepo) GetSummary(ctx context.Context, engineerID int64) (*domain.TimesheetSummary, error) {
	query := `
		SELECT COALESCE(SUM(hours_spent), 0.0), COUNT(id)
		FROM timesheet_logs
		WHERE ($1 = 0 OR engineer_id = $1)
	`
	var totalHours float64
	var count int64
	err := r.pool.QueryRow(ctx, query, engineerID).Scan(&totalHours, &count)
	if err != nil {
		return nil, err
	}

	return &domain.TimesheetSummary{
		TotalHours:     totalHours,
		EstimatedPay:   totalHours * 175000, // Basis kompensasi Rp 175.000 / jam
		TotalLogsCount: count,
	}, nil
}
