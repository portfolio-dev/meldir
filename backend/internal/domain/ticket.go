package domain

import "time"

type TicketPriority string

const (
	PriorityP1 TicketPriority = "p1_critical"
	PriorityP2 TicketPriority = "p2_major"
	PriorityP3 TicketPriority = "p3_low"
)

type TicketStatus string

const (
	TicketStatusOpen       TicketStatus = "open"
	TicketStatusInProgress TicketStatus = "in_progress"
	TicketStatusReview     TicketStatus = "review"
	TicketStatusResolved   TicketStatus = "resolved"
	TicketStatusClosed     TicketStatus = "closed"
)

type Ticket struct {
	ID                 int64          `json:"id"`
	TicketCode         string         `json:"ticket_code"`
	ProjectID          *int64         `json:"project_id,omitempty"`
	ProjectName        string         `json:"project_name,omitempty"`
	ClientID           int64          `json:"client_id"`
	ClientName         string         `json:"client_name,omitempty"`
	ClientEmail        string         `json:"client_email,omitempty"`
	ContractID         *int64         `json:"contract_id,omitempty"`
	Title              string         `json:"title"`
	Description        string         `json:"description"`
	Priority           TicketPriority `json:"priority"`
	IsMinorFeature     bool           `json:"is_minor_feature"`
	Status             TicketStatus   `json:"status"`
	AssignedEngineerID *int64         `json:"assigned_engineer_id,omitempty"`
	EngineerName       string         `json:"engineer_name,omitempty"`
	SLADeadline        time.Time      `json:"sla_deadline"`
	IsSLABreached      bool           `json:"is_sla_breached"`
	HoursSpent         float64        `json:"hours_spent"`
	ResolvedAt         *time.Time     `json:"resolved_at,omitempty"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
}

type CreateTicketRequest struct {
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	Priority       TicketPriority `json:"priority"`
	ProjectName    string         `json:"project_name,omitempty"`
	IsMinorFeature bool           `json:"is_minor_feature"`
}

type UpdateTicketStatusRequest struct {
	ID     int64        `json:"id"`
	Status TicketStatus `json:"status"`
}

type TicketFilter struct {
	Status     string
	Priority   string
	Search     string
	ClientID   int64
	EngineerID int64
}
