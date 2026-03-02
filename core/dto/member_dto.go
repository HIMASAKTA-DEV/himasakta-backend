package dto

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
)

type CreateMemberRequest struct {
	Nrp          string        `json:"nrp" binding:"required"`
	Name         string        `json:"name" binding:"required"`
	RoleId       meta.NullUUID `json:"role_id" binding:"required"`
	DepartmentId meta.NullUUID `json:"department_id"`
	PhotoId      meta.NullUUID `json:"photo_id"`
	CabinetId    meta.NullUUID `json:"cabinet_id" binding:"required"`
	Index        int           `json:"index"`
}

type UpdateMemberRequest struct {
	Nrp          *string       `json:"nrp"`
	Name         *string       `json:"name"`
	RoleId       meta.NullUUID `json:"role_id"`
	DepartmentId meta.NullUUID `json:"department_id"`
	PhotoId      meta.NullUUID `json:"photo_id"`
	CabinetId    meta.NullUUID `json:"cabinet_id"`
	Index        *int          `json:"index"`
}

type MemberGroupResponse struct {
	Role    entity.Role     `json:"role"`
	Members []entity.Member `json:"members"`
}
