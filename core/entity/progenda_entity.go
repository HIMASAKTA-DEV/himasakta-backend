package entity

import "github.com/google/uuid"

type Progenda struct {
	Timestamp
	Id           uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name         string      `gorm:"type:varchar(255);not null;unique" json:"name"`
	ThumbnailId  *uuid.UUID  `gorm:"type:uuid" json:"thumbnail_id"`
	Thumbnail    *Gallery    `gorm:"foreignKey:ThumbnailId" json:"thumbnail"`
	Goal         string      `gorm:"type:text" json:"goal"`
	Description  string      `gorm:"type:text" json:"description"`
	Timeline     string      `gorm:"type:text" json:"timeline"`
	WebsiteLink  string      `gorm:"type:varchar(255)" json:"website_link"`
	DepartmentId *uuid.UUID  `gorm:"type:uuid" json:"department_id"`
	Department   *Department `gorm:"foreignKey:DepartmentId" json:"department"`
}
