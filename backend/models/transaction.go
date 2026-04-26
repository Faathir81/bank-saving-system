package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID              string    `gorm:"primaryKey;type:uuid" json:"id"`
	AccountID       string    `gorm:"not null" json:"account_id"`
	Account         Account   `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account"`
	Type            string    `gorm:"not null" json:"type"` // "deposit" or "withdraw"
	Amount          float64   `gorm:"not null" json:"amount"`
	TransactionDate time.Time `gorm:"not null" json:"transaction_date"`
	CreatedAt       time.Time `json:"created_at"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}
