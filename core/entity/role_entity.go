package entity

import "github.com/google/uuid"

type Role struct {
	Timestamp
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Rank        string    `gorm:"type:varchar(50);not null;unique" json:"rank"`
	Index       int       `gorm:"type:int;not null" json:"index"`
	Description string    `gorm:"type:text" json:"description"`
}

func (Role) TableName() string {
	return "roles"
}
