package dto

import "github.com/google/uuid"

type CreateCabinetInfoRequest struct {
	Visi        string     `json:"visi" binding:"required"`
	Misi        string     `json:"misi" binding:"required"`
	Description string     `json:"description"`
	Tagline     string     `json:"tagline"`
	PeriodStart string     `json:"period_start" binding:"required"`
	PeriodEnd   string     `json:"period_end" binding:"required"`
	LogoId      *uuid.UUID `json:"logo_id"`
	IsActive    bool       `json:"is_active"`
}

type UpdateCabinetInfoRequest struct {
	Visi        string     `json:"visi"`
	Misi        string     `json:"misi"`
	Description string     `json:"description"`
	Tagline     string     `json:"tagline"`
	PeriodStart string     `json:"period_start"`
	PeriodEnd   string     `json:"period_end"`
	LogoId      *uuid.UUID `json:"logo_id"`
	IsActive    *bool      `json:"is_active"`
}
