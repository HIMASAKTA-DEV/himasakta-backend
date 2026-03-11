package entity

import (
	"time"

	"github.com/google/uuid"
)

type Visitor struct {
	Id         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ClientIp   string    `gorm:"type:varchar(45);not null" json:"client_ip"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:now()" json:"created_at"`
	LastSeenAt time.Time `gorm:"type:timestamp;not null;default:now()" json:"last_seen_at"`
}

func (Visitor) TableName() string {
	return "visitors"
}
