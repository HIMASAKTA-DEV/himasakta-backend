package dto

import (
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
)

type CreateTimelineRequest struct {
	ProgendaId uuid.UUID `json:"progenda_id"`
	Date       time.Time `json:"date" binding:"required"`
	Info       string    `json:"info"`
	Link       string    `json:"link"`
}
type CreateProgendaRequest struct {
	Name          string        `json:"name" binding:"required"`
	ThumbnailId   meta.NullUUID `json:"thumbnail_id"`
	Goal          string        `json:"goal"`
	Description   string        `json:"description"`
	WebsiteLink   string        `json:"website_link"`
	InstagramLink string        `json:"instagram_link"`
	TwitterLink   string        `json:"twitter_link"`
	LinkedinLink  string        `json:"linkedin_link"`
	YoutubeLink   string        `json:"youtube_link"`
	TiktokLink    string        `json:"tiktok_link"`
	
	DepartmentId  meta.NullUUID `json:"department_id"`

	Timelines []CreateTimelineRequest `json:"timelines"`
}

type UpdateTimelineRequest struct {
	Id   uuid.UUID  `json:"timeline_id"`
	Date *time.Time `json:"date"`
	Info *string    `json:"info"`
	Link *string    `json:"link"`
}
type UpdateProgendaRequest struct {
	Name          *string       `json:"name"`
	ThumbnailId   meta.NullUUID `json:"thumbnail_id"`
	Goal          *string       `json:"goal"`
	Description   *string       `json:"description"`
	WebsiteLink   *string       `json:"website_link"`
	InstagramLink *string       `json:"instagram_link"`
	TwitterLink   *string       `json:"twitter_link"`
	LinkedinLink  *string       `json:"linkedin_link"`
	YoutubeLink   *string       `json:"youtube_link"`
	TiktokLink    *string       `json:"tiktok_link"`
	DepartmentId  meta.NullUUID `json:"department_id"`

	Timelines []UpdateTimelineRequest `json:"timelines"`
}
