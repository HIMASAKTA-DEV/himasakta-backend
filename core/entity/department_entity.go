package entity

import "github.com/google/uuid"

type Department struct {
	Timestamp
	Id              uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name            string     `gorm:"type:varchar(255);not null;unique" json:"name"`
	Slug            string     `gorm:"type:varchar(255);not null;unique" json:"slug"`
	Description     string     `gorm:"type:text" json:"description"`
	LogoId          *uuid.UUID `gorm:"type:uuid;index" json:"logo_id"`
	Logo            *Gallery   `gorm:"foreignKey:LogoId;constraint:OnDelete:SET NULL" json:"logo"`
	InstagramLink   string     `gorm:"type:varchar(255)" json:"instagram_link"`
	YoutubeLink     string     `gorm:"type:varchar(255)" json:"youtube_link"`
	TwitterLink     string     `gorm:"type:varchar(255)" json:"twitter_link"`
	BankSoalLink    string     `gorm:"type:varchar(255)" json:"bank_soal_link"`
	SilabusLink     string     `gorm:"type:varchar(255)" json:"silabus_link"`
	BankRefLink     string     `gorm:"type:varchar(255)" json:"bank_ref_link"`
	LeaderId        *uuid.UUID `gorm:"type:uuid;index" json:"leader_id"`
	Leader          *Member    `gorm:"foreignKey:LeaderId;constraint:OnDelete:SET NULL" json:"leader"`

	Feeds []Gallery `gorm:"foreignKey:DepartmentId" json:"feeds"`
}

func (Department) TableName() string {
	return "departments"
}
