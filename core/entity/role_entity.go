package entity

import "github.com/google/uuid"

type Role struct {
	Timestamp
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null;unique" json:"name"` // Rank (Kahima, Kadep, etc.)
	Level       int       `gorm:"type:int;not null" json:"level"`                // Index for sorting roles
	Description string    `gorm:"type:text" json:"description"`
}

func (Role) TableName() string {
	return "roles"
}
