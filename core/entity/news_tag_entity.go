package entity

import "github.com/google/uuid"

type NewsTag struct {
	NewsId uuid.UUID `gorm:"type:uuid;primaryKey;constraint:OnDelete:SET NULL" json:"news_id"`
	TagId  uuid.UUID `gorm:"type:uuid;primaryKey;constraint:OnDelete:SET NULL" json:"tag_id"`
}

func (NewsTag) TablesName() string {
	return "news_tags"
}
