package dto

import "github.com/google/uuid"

type CreateCabinetInfoRequest struct {
	Visi    		string     `json:"visi" binding:"required"`
	Misi     		string     `json:"misi" binding:"required"`
	Tagline  		string     `json:"tagline"`
	Period   		string     `json:"period" binding:"required"`
	LogoId   		*uuid.UUID `json:"logo_id"`
	OrganigramId    *uuid.UUID `json:"organigram_id"`
	IsActive 		bool       `json:"is_active"`
}

type UpdateCabinetInfoRequest struct {
	Visi     		string     `json:"visi"`
	Misi     		string     `json:"misi"`
	Tagline  		string     `json:"tagline"`
	Period   		string     `json:"period"`
	LogoId   		*uuid.UUID `json:"logo_id"`
	OrganigramId    *uuid.UUID `json:"organigram_id"`
	IsActive 		*bool       `json:"is_active"`
}
