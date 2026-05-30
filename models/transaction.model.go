package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"` // "credit" or "debit" other "income" or "expense"
	Category  string    `json:"category"`
	Note      string    `json:"note"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}
