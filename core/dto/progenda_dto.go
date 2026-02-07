package dto

import "github.com/google/uuid"

type CreateProgendaRequest struct {
	Name         string     `json:"name" binding:"required"`
	ThumbnailId  *uuid.UUID `json:"thumbnail_id"`
	Goal         string     `json:"goal"`
	Description  string     `json:"description"`
	Timeline     string     `json:"timeline"`
	WebsiteLink  string     `json:"website_link"`
	DepartmentId *uuid.UUID `json:"department_id"`
}

type UpdateProgendaRequest struct {
	Name         string     `json:"name"`
	ThumbnailId  *uuid.UUID `json:"thumbnail_id"`
	Goal         string     `json:"goal"`
	Description  string     `json:"description"`
	Timeline     string     `json:"timeline"`
	WebsiteLink  string     `json:"website_link"`
	DepartmentId *uuid.UUID `json:"department_id"`
}
