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
	Content     string     `gorm:"type:text;not null" json:"content"`
	ThumbnailId *uuid.UUID `gorm:"type:uuid;index;constraint:OnDelete:SET NULL" json:"thumbnail_id"`
	Thumbnail   *Gallery   `gorm:"foreignKey:ThumbnailId" json:"thumbnail"`
	PublishedAt time.Time  `gorm:"type:timestamp" json:"published_at"`
	AuthorId    *uuid.UUID `gorm:"type:uuid;index;constraint:OnDelete:SET NULL" json:"author_id"`
	Author      *Member    `gorm:"foreignKey:AuthorId" json:"author"`

	Hashtags []Tag `gorm:"many2many:news_tags" json:"tags"`
}

func (News) TableName() string {
	return "news"
}
