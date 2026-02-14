package entity

import (
	"time"

	"github.com/google/uuid"
)

type Timeline struct {
	Timestamp
	Id         uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProgendaId *uuid.UUID `gorm:"type:uuid" json:"progenda_id"`
	Date       time.Time  `gorm:"type:date" json:"date"`
	Info       string     `gorm:"type:varchar(100)" json:"info"`
	Link       string     `gorm:"type:varchar(100)" json:"link"`
}

func (Timeline) TablesName() string {
	return "timelines"
}
