package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateNewsRequest struct {
	Title       string     `json:"title" binding:"required"`
	Tagline     string     `json:"tagline"`
	Hashtags    string     `json:"hashtags"`
	Content     string     `json:"content" binding:"required"`
	ThumbnailId *uuid.UUID `json:"thumbnail_id"`
	PublishedAt time.Time  `json:"published_at"`
}

type UpdateNewsRequest struct {
	Title       string     `json:"title"`
	Tagline     string     `json:"tagline"`
	Hashtags    string     `json:"hashtags"`
	Content     string     `json:"content"`
	ThumbnailId *uuid.UUID `json:"thumbnail_id"`
	PublishedAt *time.Time `json:"published_at"`
}
