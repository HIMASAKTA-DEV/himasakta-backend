package entity

import "github.com/google/uuid"

type Member struct {
	Timestamp
	Id           uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Nrp          string      `gorm:"type:varchar(20);not null" json:"nrp"`
	Name         string      `gorm:"type:varchar(255);not null;unique" json:"name"`
	Role         string      `gorm:"type:varchar(100);not null" json:"role"`
	DepartmentId *uuid.UUID  `gorm:"type:uuid" json:"department_id"`
	Department   *Department `gorm:"foreignKey:DepartmentId" json:"department"`
	PhotoId      *uuid.UUID  `gorm:"type:uuid" json:"photo_id"`
	Photo        *Gallery    `gorm:"foreignKey:PhotoId" json:"photo"`
	Period       string      `gorm:"type:varchar(10);not null" json:"period"`
}
