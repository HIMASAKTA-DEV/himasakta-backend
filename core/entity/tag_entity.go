package entity

import "github.com/google/uuid"

type Tag struct {
	Id   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"type:varchar(255);unique" json:"name"`
}

func (Tag) TablesName() string {
	return "tags"
}
