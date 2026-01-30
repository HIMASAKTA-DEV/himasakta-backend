package entity

import (
	"github.com/google/uuid"
)

type Item struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CategoryCode string    `json:"category_code" gorm:"not null"`
	CategoryName string    `json:"category_name" gorm:"not null"`
	Group        string    `json:"group" gorm:"not null"`
	AccountCode  string    `json:"account_code" gorm:"not null"`
	Description  string    `json:"description" gorm:"not null"`

	Timestamp
}

func (i *Item) TableName() string {
	return "items"
}

