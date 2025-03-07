package entities

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID `json:"id"`
	TaxID         string    `json:"tax_id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	MonthlyIncome int64     `json:"monthly_income,omitempty"`
	AnnualRevenue int64     `json:"annual_revenue,omitempty"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at,omitempty"`
}
