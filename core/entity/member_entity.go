package entity

import "github.com/google/uuid"

type Member struct {
	Timestamp
	Id           uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Nrp          string       `gorm:"type:varchar(20);not null" json:"nrp"`
	Name         string       `gorm:"type:varchar(255);not null;unique" json:"name"`
	RoleId       *uuid.UUID   `gorm:"type:uuid" json:"role_id"`
	Role         *Role        `gorm:"foreignKey:RoleId" json:"role"`
	Index        int          `gorm:"type:int;not null" json:"index"`
	DepartmentId *uuid.UUID   `gorm:"type:uuid" json:"department_id"`
	Department   *Department  `gorm:"foreignKey:DepartmentId" json:"department"`
	PhotoId      *uuid.UUID   `gorm:"type:uuid" json:"photo_id"`
	Photo        *Gallery     `gorm:"foreignKey:PhotoId" json:"photo"`
	CabinetId    *uuid.UUID   `gorm:"type:uuid" json:"cabinet_id"`
	Cabinet      *CabinetInfo `gorm:"foreignKey:CabinetId" json:"cabinet"`
}

func (Member) TableName() string {
	return "members"
}
