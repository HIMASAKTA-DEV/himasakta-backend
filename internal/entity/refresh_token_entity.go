package entity

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID           uuid.UUID `json:"id" gorm:"not null;type:uuid;default:uuid_generate_v4()"`
	UserID       uuid.UUID `json:"user_id" gorm:"not null"`
	RefreshToken string    `json:"refresh_token" gorm:"not null;uniqueIndex"`
	UserAgent    string    `json:"user_agent" gorm:"not null"`
	IP           string    `json:"ip" gorm:"not null"`
	ExpiresAt    time.Time `json:"expires_at" gorm:"not null"`

	Timestamp
}
