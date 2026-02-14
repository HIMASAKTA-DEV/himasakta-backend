package entity

import (
	"time"

	"github.com/google/uuid"
)

type Progenda struct {
	Timestamp
	Id            uuid.UUID          `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name          string             `gorm:"type:varchar(255);not null;unique" json:"name"`
	ThumbnailId   *uuid.UUID         `gorm:"type:uuid" json:"thumbnail_id"`
	Thumbnail     *Gallery           `gorm:"foreignKey:ThumbnailId" json:"thumbnail"`
	Goal          string             `gorm:"type:text" json:"goal"`
	Description   string             `gorm:"type:text" json:"description"`
	InstagramLink string             `gorm:"type:varchar(255)" json:"instagram_link"`
	TwitterLink   string             `gorm:"type:varchar(255)" json:"twitter_link"`
	YoutubeLink   string             `gorm:"type:varchar(255)" json:"youtube_link"`
	LinkedinLink  string             `gorm:"type:varchar(255)" json:"linkedin_link"`
	WebsiteLink   string             `gorm:"type:varchar(255)" json:"website_link"`
	DepartmentId  *uuid.UUID         `gorm:"type:uuid" json:"department_id"`
	Department    *Department        `gorm:"foreignKey:DepartmentId" json:"department"`
	Timelines     []ProgendaTimeline `gorm:"foreignKey:ProgendaId;constraint:OnDelete:CASCADE" json:"timelines"`
}

type ProgendaTimeline struct {
	Timestamp
	Id         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProgendaId uuid.UUID `gorm:"type:uuid;not null" json:"progenda_id"`
	EventName  string    `gorm:"type:varchar(255);not null" json:"event_name"`
	Date       time.Time `gorm:"type:timestamp without time zone;not null" json:"date"`
}

func (Progenda) TableName() string {
	return "progendas"
}

func (ProgendaTimeline) TableName() string {
	return "progenda_timelines"
}
