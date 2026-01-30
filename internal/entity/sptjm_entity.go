package entity

import (
	"github.com/google/uuid"
)

type SPTJM struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	SKPDName  string    `json:"skpd_name" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	PIC1Name  string    `json:"pic1_name" gorm:"not null"`
	PIC1Email string    `json:"pic1_email" gorm:"not null"`
	PIC1Phone string    `json:"pic1_phone" gorm:"not null"`
	PIC2Name  string    `json:"pic2_name" gorm:"not null"`
	PIC2Email string    `json:"pic2_email" gorm:"not null"`
	PIC2Phone string    `json:"pic2_phone" gorm:"not null"`
	FileURL   string    `json:"file_url" gorm:"not null"`

	Timestamp
}

func (s *SPTJM) TableName() string {
	return "sptjms"
}
