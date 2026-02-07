package dto

import "github.com/google/uuid"

type CreateMemberRequest struct {
	Nrp          string     `json:"nrp" binding:"required"`
	Name         string     `json:"name" binding:"required"`
	Role         string     `json:"role" binding:"required"`
	DepartmentId *uuid.UUID `json:"department_id"`
	PhotoId      *uuid.UUID `json:"photo_id"`
	Period       string     `json:"period" binding:"required"`
}

type UpdateMemberRequest struct {
	Nrp          string     `json:"nrp"`
	Name         string     `json:"name"`
	Role         string     `json:"role"`
	DepartmentId *uuid.UUID `json:"department_id"`
	PhotoId      *uuid.UUID `json:"photo_id"`
	Period       string     `json:"period"`
}
