package domain

import "time"

type InvoiceStatus string

const (
	InvoiceStatusUnpaid    InvoiceStatus = "unpaid"
	InvoiceStatusVerifying InvoiceStatus = "verifying"
	InvoiceStatusPaid      InvoiceStatus = "paid"
	InvoiceStatusOverdue   InvoiceStatus = "overdue"
	InvoiceStatusCancelled InvoiceStatus = "cancelled"
)

type Invoice struct {
	ID               int64         `json:"id"`
	InvoiceNumber    string        `json:"invoice_number"`
	ClientID         int64         `json:"client_id"`
	ClientName       string        `json:"client_name"`
	ContractID       *int64        `json:"contract_id,omitempty"`
	Amount           float64       `json:"amount"`       // DPP (Dasar Pengenaan Pajak)
	TaxAmount        float64       `json:"tax_amount"`   // PPN 11%
	TotalAmount      float64       `json:"total_amount"` // DPP + PPN
	DueDate          string        `json:"due_date"`
	Status           InvoiceStatus `json:"status"`
	BankDestination  string        `json:"bank_destination"`
	TaxInvoiceNumber *string       `json:"tax_invoice_number,omitempty"`
	PaidAt           *time.Time    `json:"paid_at,omitempty"`
	CreatedAt        time.Time     `json:"created_at"`
}

type CreateInvoiceRequest struct {
	ClientID    int64   `json:"client_id"`
	ClientName  string  `json:"client_name"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"` // DPP
	DueDate     string  `json:"due_date"`
}

type UpdateInvoiceStatusRequest struct {
	ID     int64         `json:"id"`
	Status InvoiceStatus `json:"status"`
}
