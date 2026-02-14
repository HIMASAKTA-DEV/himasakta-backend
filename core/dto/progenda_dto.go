package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateTimelineRequest struct {
	ProgendaId uuid.UUID `json:"progenda_id" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	Info       string    `json:"info"`
	Link       string    `json:"link"`
}
type CreateProgendaRequest struct {
	Name          string     `json:"name" binding:"required"`
	ThumbnailId   *uuid.UUID `json:"thumbnail_id"`
	Goal          string     `json:"goal"`
	Description   string     `json:"description"`
	WebsiteLink   string     `json:"website_link"`
	InstagramLink string     `json:"instagram_link"`
	TwitterLink   string     `json:"twitter_link"`
	LinkedinLink  string     `json:"linkedin_link"`
	YoutubeLink   string     `json:"youtube_link"`
	DepartmentId  *uuid.UUID `json:"department_id"`

	Timelines []CreateTimelineRequest `json:"timelines"`
}

type UpdateTimelineRequest struct {
	Id   uuid.UUID `json:"timeline_id"`
	Date time.Time `json:"date" binding:"required"`
	Info string    `json:"info"`
	Link string    `json:"link"`
}
type UpdateProgendaRequest struct {
	Name          string     `json:"name"`
	ThumbnailId   *uuid.UUID `json:"thumbnail_id"`
	Goal          string     `json:"goal"`
	Description   string     `json:"description"`
	WebsiteLink   string     `json:"website_link"`
	InstagramLink string     `json:"instagram_link"`
	TwitterLink   string     `json:"twitter_link"`
	LinkedinLink  string     `json:"linkedin_link"`
	YoutubeLink   string     `json:"youtube_link"`
	DepartmentId  *uuid.UUID `json:"department_id"`

	Timelines []UpdateTimelineRequest `json:"timelines"`
}
