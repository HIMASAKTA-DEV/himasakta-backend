package dto

import "github.com/google/uuid"

type CreateMemberRequest struct {
	Nrp          string     `json:"nrp" binding:"required"`
	Name         string     `json:"name" binding:"required"`
	RoleId       *uuid.UUID `json:"role_id" binding:"required"`
	DepartmentId *uuid.UUID `json:"department_id"`
	PhotoId      *uuid.UUID `json:"photo_id"`
	CabinetId    *uuid.UUID `json:"cabinet_id" binding:"required"`
	Index        int        `json:"index"`
}

type UpdateMemberRequest struct {
	Nrp          string     `json:"nrp"`
	Name         string     `json:"name"`
	RoleId       *uuid.UUID `json:"role_id"`
	DepartmentId *uuid.UUID `json:"department_id"`
	PhotoId      *uuid.UUID `json:"photo_id"`
	CabinetId    *uuid.UUID `json:"cabinet_id"`
	Index        int        `json:"index"`
}
