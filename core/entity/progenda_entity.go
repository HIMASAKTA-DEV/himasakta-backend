package entity

import "github.com/google/uuid"

type Progenda struct {
	Timestamp
	Id            uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name          string      `gorm:"type:varchar(255);not null;unique" json:"name"`
	ThumbnailId   *uuid.UUID  `gorm:"type:uuid" json:"thumbnail_id"`
	Thumbnail     *Gallery    `gorm:"foreignKey:ThumbnailId" json:"thumbnail"`
	Goal          string      `gorm:"type:text" json:"goal"`
	Description   string      `gorm:"type:text" json:"description"`
	WebsiteLink   string      `gorm:"type:varchar(255)" json:"website_link"`
	InstagramLink string      `gorm:"type:varchar(255)" json:"instagram_link"`
	TwitterLink   string      `gorm:"type:varchar(255)" json:"twitter_link"`
	LinkedinLink  string      `gorm:"type:varchar(255)" json:"linkedin_link"`
	YoutubeLink   string      `gorm:"type:varchar(255)" json:"youtube_link"`
	DepartmentId  *uuid.UUID  `gorm:"type:uuid" json:"department_id"`
	Department    *Department `gorm:"foreignKey:DepartmentId" json:"department"`

	Timelines []Timeline `gorm:"foreignKey:ProgendaId" json:"timelines"`
	Feeds     []Gallery  `gorm:"foreignKey:ProgendaId" json:"feeds"`
}

func (Progenda) TableName() string {
	return "progendas"
}
