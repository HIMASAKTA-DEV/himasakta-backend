package dto

import "github.com/google/uuid"

type CreateDepartmentRequest struct {
	Name            string     `json:"name" binding:"required"`
	Description     string     `json:"description"`
	LogoId          *uuid.UUID `json:"logo_id"`
	SocialMediaLink string     `json:"social_media_link"`
	BankSoalLink    string     `json:"bank_soal_link"`
	SilabusLink     string     `json:"silabus_link"`
}

type UpdateDepartmentRequest struct {
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	LogoId          *uuid.UUID `json:"logo_id"`
	SocialMediaLink string     `json:"social_media_link"`
	BankSoalLink    string     `json:"bank_soal_link"`
	SilabusLink     string     `json:"silabus_link"`
}
