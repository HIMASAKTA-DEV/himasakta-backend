package dto

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
)

type CreateDepartmentRequest struct {
	Name            string        `json:"name" binding:"required"`
	Description     string        `json:"description"`
	LogoId          meta.NullUUID `json:"logo_id"`
	InstagramLink   string        `json:"instagram_link"`
	YoutubeLink     string        `json:"youtube_link"`
	TwitterLink     string        `json:"twitter_link"`
	BankSoalLink    string        `json:"bank_soal_link"`
	SilabusLink     string        `json:"silabus_link"`
	BankRefLink     string        `json:"bank_ref_link"`
	LeaderId        meta.NullUUID `json:"leader_id"`
}

type UpdateDepartmentRequest struct {
	Name            *string       `json:"name"`
	Description     *string       `json:"description"`
	LogoId          meta.NullUUID `json:"logo_id"`
	InstagramLink   *string       `json:"instagram_link"`
	YoutubeLink     *string       `json:"youtube_link"`
	TwitterLink     *string       `json:"twitter_link"`
	BankSoalLink    *string       `json:"bank_soal_link"`
	SilabusLink     *string       `json:"silabus_link"`
	BankRefLink     *string       `json:"bank_ref_link"`
	LeaderId        meta.NullUUID `json:"leader_id"`
}
