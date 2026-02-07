package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateMonthlyEventRequest struct {
	Title       string     `json:"title" binding:"required"`
	ThumbnailId *uuid.UUID `json:"thumbnail_id"`
	Description string     `json:"description"`
	Month       time.Time  `json:"month" binding:"required"`
	Link        string     `json:"link"`
}

type UpdateMonthlyEventRequest struct {
	Title       string     `json:"title"`
	ThumbnailId *uuid.UUID `json:"thumbnail_id"`
	Description string     `json:"description"`
	Month       *time.Time `json:"month"`
	Link        string     `json:"link"`
}
