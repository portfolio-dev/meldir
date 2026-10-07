package domain

import "time"

type TimesheetLog struct {
	ID              int64     `json:"id"`
	EngineerID      int64     `json:"engineer_id"`
	EngineerName    string    `json:"engineer_name,omitempty"`
	ProjectID       *int64    `json:"project_id,omitempty"`
	ProjectName     string    `json:"project_name,omitempty"`
	TicketID        *int64    `json:"ticket_id,omitempty"`
	TicketCode      string    `json:"ticket_code,omitempty"`
	AddonOrderID    *int64    `json:"addon_order_id,omitempty"`
	HoursSpent      float64   `json:"hours_spent"`
	WorkDescription string    `json:"work_description"`
	LogDate         string    `json:"log_date"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateTimesheetRequest struct {
	TicketCode      string  `json:"ticket_code,omitempty"`
	ProjectName     string  `json:"project_name"`
	HoursSpent      float64 `json:"hours_spent"`
	WorkDescription string  `json:"work_description"`
	LogDate         string  `json:"log_date"`
}

type TimesheetSummary struct {
	TotalHours     float64 `json:"total_hours"`
	EstimatedPay   float64 `json:"estimated_pay"`
	TotalLogsCount int64   `json:"total_logs_count"`
}
