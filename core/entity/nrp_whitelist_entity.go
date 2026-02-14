package entity

import "github.com/google/uuid"

type NrpWhitelist struct {
	Timestamp
	Id   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Nrp  string    `gorm:"type:varchar(255);not null;unique"`
	Name string    `gorm:"type:varchar(255);not null"`
}

func (NrpWhitelist) TableName() string {
	return "nrpWhitelists"
}
