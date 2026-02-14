package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProgendaTimelineRequest struct {
	EventName string    `json:"event_name" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
}

type CreateProgendaRequest struct {
	Name          string                    `json:"name" binding:"required"`
	ThumbnailId   *uuid.UUID                `json:"thumbnail_id"`
	Goal          string                    `json:"goal"`
	Description   string                    `json:"description"`
	InstagramLink string                    `json:"instagram_link"`
	TwitterLink   string                    `json:"twitter_link"`
	YoutubeLink   string                    `json:"youtube_link"`
	LinkedinLink  string                    `json:"linkedin_link"`
	WebsiteLink   string                    `json:"website_link"`
	DepartmentId  *uuid.UUID                `json:"department_id"`
	Timelines     []ProgendaTimelineRequest `json:"timelines"`
}

type UpdateProgendaRequest struct {
	Name          string                    `json:"name"`
	ThumbnailId   *uuid.UUID                `json:"thumbnail_id"`
	Goal          string                    `json:"goal"`
	Description   string                    `json:"description"`
	InstagramLink string                    `json:"instagram_link"`
	TwitterLink   string                    `json:"twitter_link"`
	YoutubeLink   string                    `json:"youtube_link"`
	LinkedinLink  string                    `json:"linkedin_link"`
	WebsiteLink   string                    `json:"website_link"`
	DepartmentId  *uuid.UUID                `json:"department_id"`
	Timelines     []ProgendaTimelineRequest `json:"timelines"`
}
