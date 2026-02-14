package dto

import "github.com/google/uuid"

type CreateCabinetInfoRequest struct {
	Visi         string     `json:"visi" binding:"required"`
	Misi         string     `json:"misi" binding:"required"`
	Description  string     `json:"description"`
	Tagline      string     `json:"tagline"`
	PeriodStart  string     `json:"period_start" binding:"required"`
	PeriodEnd    string     `json:"period_end" binding:"required"`
	LogoId       *uuid.UUID `json:"logo_id"`
	OrganigramId *uuid.UUID `json:"organigram_id"`
	IsActive     bool       `json:"is_active"`
}

type UpdateCabinetInfoRequest struct {
	Visi         string     `json:"visi" binding:"required"`
	Misi         string     `json:"misi" binding:"required"`
	Description  string     `json:"description"`
	Tagline      string     `json:"tagline"`
	PeriodStart  string     `json:"period_start" binding:"required"`
	PeriodEnd    string     `json:"period_end" binding:"required"`
	LogoId       *uuid.UUID `json:"logo_id"`
	OrganigramId *uuid.UUID `json:"organigram_id"`
	IsActive     *bool      `json:"is_active"`
}
