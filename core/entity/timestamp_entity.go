package entity

import (
	"time"
)

type Timestamp struct {
	CreatedAt time.Time `gorm:"type:timestamp without time zone" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp without time zone" json:"updated_at"`
}
