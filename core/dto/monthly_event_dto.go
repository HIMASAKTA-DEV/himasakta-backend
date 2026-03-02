package dto

import (
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
)

type CreateMonthlyEventRequest struct {
	Title       string        `json:"title" binding:"required"`
	ThumbnailId meta.NullUUID `json:"thumbnail_id"`
	Description string        `json:"description"`
	Month       time.Time     `json:"month" binding:"required"`
	Link        string        `json:"link"`
}

type UpdateMonthlyEventRequest struct {
	Title       *string       `json:"title"`
	ThumbnailId meta.NullUUID `json:"thumbnail_id"`
	Description *string       `json:"description"`
	Month       *time.Time    `json:"month"`
	Link        *string       `json:"link"`
}
