package entity

import "github.com/google/uuid"

type Gallery struct {
	Timestamp
	Id           uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ImageUrl     string     `gorm:"type:varchar(255);not null" json:"image_url"`
	Caption      string     `gorm:"type:varchar(255)" json:"caption"`
	Category     string     `gorm:"type:varchar(100)" json:"category"`
	DepartmentId *uuid.UUID `gorm:"type:uuid" json:"department_id"`

	ProgendaId *uuid.UUID `gorm:"type:uuid" json:"progenda_id"`
}

func (Gallery) TableName() string {
	return "galleries"
}
