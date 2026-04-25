package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID             string       `gorm:"primaryKey;type:uuid" json:"id"`
	CustomerID     string       `gorm:"not null" json:"customer_id"`
	Customer       Customer     `gorm:"foreignKey:CustomerID" json:"customer"`
	DepositoTypeID string       `gorm:"not null" json:"deposito_type_id"`
	DepositoType   DepositoType `gorm:"foreignKey:DepositoTypeID" json:"deposito_type"`
	Balance        float64      `gorm:"default:0" json:"balance"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return
}
