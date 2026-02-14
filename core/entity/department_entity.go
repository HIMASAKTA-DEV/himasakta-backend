package entity

import "github.com/google/uuid"

type Department struct {
	Timestamp
	Id              uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name            string     `gorm:"type:varchar(255);not null;unique" json:"name"`
	Description     string     `gorm:"type:text" json:"description"`
	LogoId          *uuid.UUID `gorm:"type:uuid" json:"logo_id"`
	Logo            *Gallery   `gorm:"foreignKey:LogoId" json:"logo"`
	SocialMediaLink string     `gorm:"type:varchar(255)" json:"social_media_link"`
	BankSoalLink    string     `gorm:"type:varchar(255)" json:"bank_soal_link"`
	SilabusLink     string     `gorm:"type:varchar(255)" json:"silabus_link"`
	BankRefLink     string     `gorm:"type:varchar(255)" json:"bank_ref_link"`
}

func (Department) TableName() string {
	return "departments"
}
