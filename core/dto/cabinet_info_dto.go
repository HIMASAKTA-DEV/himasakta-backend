package dto

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
)

type CreateCabinetInfoRequest struct {
	Visi         string        `json:"visi" binding:"required"`
	Misi         string        `json:"misi" binding:"required"`
	Description  string        `json:"description"`
	Tagline      string        `json:"tagline"`
	PeriodStart  string        `json:"period_start" binding:"required"`
	PeriodEnd    string        `json:"period_end" binding:"required"`
	LogoId       meta.NullUUID `json:"logo_id"`
	OrganigramId meta.NullUUID `json:"organigram_id"`
	IsActive     *bool         `json:"is_active"`
}

type UpdateCabinetInfoRequest struct {
	Visi         *string       `json:"visi"`
	Misi         *string       `json:"misi"`
	Description  *string       `json:"description"`
	Tagline      *string       `json:"tagline"`
	PeriodStart  *string       `json:"period_start"`
	PeriodEnd    *string       `json:"period_end"`
	LogoId       meta.NullUUID `json:"logo_id"`
	OrganigramId meta.NullUUID `json:"organigram_id"`
	IsActive     *bool         `json:"is_active"`
}
