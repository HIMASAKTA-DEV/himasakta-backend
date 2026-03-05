package dto

import (
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
)

type CreateNewsRequest struct {
	Title       string        `json:"title" binding:"required"`
	Tagline     string        `json:"tagline"`
	Hashtags    string        `json:"hashtags"`
	Content     string        `json:"content" binding:"required"`
	ThumbnailId meta.NullUUID `json:"thumbnail_id"`
	PublishedAt time.Time     `json:"published_at"`
	AuthorId    meta.NullUUID `json:"author_id"`
}

type UpdateNewsRequest struct {
	Title       *string       `json:"title"`
	Tagline     *string       `json:"tagline"`
	Hashtags    *string       `json:"hashtags"`
	Content     *string       `json:"content"`
	ThumbnailId meta.NullUUID `json:"thumbnail_id"`
	PublishedAt *time.Time    `json:"published_at"`
	AuthorId    meta.NullUUID `json:"author_id"`
}
