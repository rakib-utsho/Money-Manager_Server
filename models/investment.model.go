package models

import "time"

// Investment represents an investment belonging to a user.
type Investment struct {
	ID            int64      `json:"id"`
	UserID        int64      `json:"user_id"`
	Type          string     `json:"type"`
	Amount        float64    `json:"amount"`
	MonthlyAmount *float64   `json:"monthly_amount"`
	InterestRate  *float64   `json:"interest_rate"`
	Duration      *int       `json:"duration"` // in months
	MaturityDate  *time.Time `json:"maturity_date"`
	CurrentValue  *float64   `json:"current_value"`
	Notes         string     `json:"notes"`
	CreatedAt     time.Time  `json:"created_at"`
}
