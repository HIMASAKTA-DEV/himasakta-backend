package entity

import (
	"time"

	"github.com/google/uuid"
)

type NrpWhitelist struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Nrp       string    `gorm:"type:varchar(255);not null;unique"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp without time zone" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp without time zone" json:"updated_at"`
}
