package domain

import "time"

type LeadStatus string

const (
	LeadStatusNew          LeadStatus = "new"
	LeadStatusContacted    LeadStatus = "contacted"
	LeadStatusQuoted       LeadStatus = "quoted"
	LeadStatusWonContract  LeadStatus = "won_contract"
	LeadStatusLost         LeadStatus = "lost"
)

type Lead struct {
	ID              int64      `json:"id"`
	LeadCode        string     `json:"lead_code"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	PhoneWA         string     `json:"phone_wa"`
	CompanyName     string     `json:"company_name"`
	ServiceInterest string     `json:"service_interest"`
	BudgetRange     string     `json:"budget_range"`
	Message         string     `json:"message"`
	Status          LeadStatus `json:"status"`
	Source          string     `json:"source"`
	AssignedAdminID *int64     `json:"assigned_admin_id,omitempty"`
	Notes           string     `json:"notes"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type CreateLeadRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	PhoneWA         string `json:"phone_wa"`
	CompanyName     string `json:"company_name"`
	ServiceInterest string `json:"service_interest"`
	BudgetRange     string `json:"budget_range"`
	Message         string `json:"message"`
	Source          string `json:"source"`
}

type UpdateLeadStatusRequest struct {
	ID     int64      `json:"id"`
	Status LeadStatus `json:"status"`
	Notes  string     `json:"notes"`
}
