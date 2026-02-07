package entity

import (
	"time"

	"github.com/google/uuid"
)

type MonthlyEvent struct {
	Timestamp
	Id          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Title       string     `gorm:"type:varchar(255);not null;unique" json:"title"`
	ThumbnailId *uuid.UUID `gorm:"type:uuid" json:"thumbnail_id"`
	Thumbnail   *Gallery   `gorm:"foreignKey:ThumbnailId" json:"thumbnail"`
	Description string     `gorm:"type:text" json:"description"`
	Month       time.Time  `gorm:"type:date;not null" json:"month"`
	Link        string     `gorm:"type:varchar(255)" json:"link"`
}
