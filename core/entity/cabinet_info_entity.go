package entity

import "github.com/google/uuid"

type CabinetInfo struct {
	Timestamp
	Id          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Visi        string     `gorm:"type:text;not null" json:"visi"`
	Misi        string     `gorm:"type:text;not null" json:"misi"`
	Description string     `gorm:"type:text" json:"description"`
	Tagline     string     `gorm:"type:varchar(255)" json:"tagline"`
	PeriodStart string     `gorm:"type:date;not null" json:"period_start"`
	PeriodEnd   string     `gorm:"type:date;not null" json:"period_end"`
	LogoId      *uuid.UUID `gorm:"type:uuid" json:"logo_id"`
	Logo        *Gallery   `gorm:"foreignKey:LogoId" json:"logo"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`
}

func (CabinetInfo) TableName() string {
	return "cabinet_infos"
}
