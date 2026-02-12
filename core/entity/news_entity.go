package entity

import (
	"time"

	"github.com/google/uuid"
)

type News struct {
	Timestamp
	Id          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Title       string     `gorm:"type:varchar(255);not null;unique" json:"title"`
	Slug        string     `gorm:"type:varchar(255);not null;unique" json:"slug"`
	Tagline     string     `gorm:"type:varchar(255)" json:"tagline"`
	Hashtags    string     `gorm:"type:text" json:"hashtags"` // comma separated
	Content     string     `gorm:"type:text;not null" json:"content"`
	ThumbnailId *uuid.UUID `gorm:"type:uuid" json:"thumbnail_id"`
	Thumbnail   *Gallery   `gorm:"foreignKey:ThumbnailId" json:"thumbnail"`
	PublishedAt time.Time  `gorm:"type:timestamp" json:"published_at"`
}

func (News) TableName() string {
	return "news"
}
