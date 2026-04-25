package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepositoType struct {
	ID           string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	YearlyReturn float64   `gorm:"not null" json:"yearly_return"` // 0.03 for 3%
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (dt *DepositoType) BeforeCreate(tx *gorm.DB) (err error) {
	dt.ID = uuid.New().String()
	return
}
