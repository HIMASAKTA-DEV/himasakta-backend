package entity

import "github.com/google/uuid"

type CabinetInfo struct {
	Timestamp
	Id       		uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Visi     		string     `gorm:"type:text;not null" json:"visi"`
	Misi     		string     `gorm:"type:text;not null" json:"misi"`
	Tagline  		string     `gorm:"type:varchar(255)" json:"tagline"`
	Period   		string     `gorm:"type:varchar(10);not null" json:"period"`
	LogoId   		*uuid.UUID `gorm:"type:uuid" json:"logo_id"`
	Logo     		*Gallery   `gorm:"foreignKey:LogoId" json:"logo"`
	OrganigramId    *uuid.UUID `gorm:"type:uuid" json:"organigram_id"`
	Organigram      *Gallery   `gorm:"foreignKey:OrganigramId" json:"organigram"`
	IsActive 		bool       `gorm:"default:true" json:"is_active"`
}
