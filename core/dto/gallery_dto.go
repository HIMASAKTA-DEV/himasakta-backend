package dto

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
)

type CreateGalleryRequest struct {
	ImageUrl     string        `json:"image_url" binding:"required"`
	Caption      string        `json:"caption"`
	Category     string        `json:"category"`
	DepartmentId meta.NullUUID `json:"department_id"`
	ProgendaId   meta.NullUUID `json:"progenda_id"`
	CabinetId    meta.NullUUID `json:"cabinet_id"`
}

type UpdateGalleryRequest struct {
	ImageUrl     *string       `json:"image_url"`
	Caption      *string       `json:"caption"`
	Category     *string       `json:"category"`
	DepartmentId meta.NullUUID `json:"department_id"`
	ProgendaId   meta.NullUUID `json:"progenda_id"`
	CabinetId    meta.NullUUID `json:"cabinet_id"`
}
