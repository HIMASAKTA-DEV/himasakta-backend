package dto

import "github.com/google/uuid"

type CreateGalleryRequest struct {
	ImageUrl     string     `json:"image_url" binding:"required"`
	Caption      string     `json:"caption"`
	Category     string     `json:"category"`
	DepartmentId *uuid.UUID `json:"department_id"`
}

type UpdateGalleryRequest struct {
	ImageUrl     string     `json:"image_url"`
	Caption      string     `json:"caption"`
	Category     string     `json:"category"`
	DepartmentId *uuid.UUID `json:"department_id"`
}
