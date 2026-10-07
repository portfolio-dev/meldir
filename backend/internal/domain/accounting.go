package domain

import "time"

type JournalEntry struct {
	ID                int64     `json:"id"`
	JournalNumber     string    `json:"journal_number"`
	JournalDate       string    `json:"journal_date"`
	SourceType        string    `json:"source_type"`
	SourceReferenceID string    `json:"source_reference_id,omitempty"`
	Memo              string    `json:"memo"`
	DebitAccount      string    `json:"debit_account"`
	CreditAccount     string    `json:"credit_account"`
	Amount            float64   `json:"amount"`
	TotalDebit        float64   `json:"total_debit"`
	TotalCredit       float64   `json:"total_credit"`
	IsPosted          bool      `json:"is_posted"`
	CreatedBy         int64     `json:"created_by"`
	CreatedAt         time.Time `json:"created_at"`
}

type CreateJournalRequest struct {
	Date          string  `json:"date"`
	Description   string  `json:"description"`
	DebitAccount  string  `json:"debit_account"`
	CreditAccount string  `json:"credit_account"`
	Amount        float64 `json:"amount"`
}

type COAAccount struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Balance  string  `json:"balance"`
	Amount   float64 `json:"amount"`
}
