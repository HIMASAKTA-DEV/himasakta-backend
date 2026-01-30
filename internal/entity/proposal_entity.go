package entity

import (
	"github.com/google/uuid"
)

type Proposal struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	No            string    `json:"no" gorm:"not null"`
	ComponentName string    `json:"component_name" gorm:"not null"`
	Specification string    `json:"specification" gorm:"not null"`
	Unit          string    `json:"unit" gorm:"not null"`
	UnitPrice     float64   `json:"unit_price" gorm:"not null"`
	AccountCode   string    `json:"account_code" gorm:"not null"`

	Timestamp
}

func (p *Proposal) TableName() string {
	return "proposals"
}

