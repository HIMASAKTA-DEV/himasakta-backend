package dto

import (
	"time"

	"github.com/google/uuid"
)

type AcademicCalendarItem struct {
	Id           uuid.UUID  `json:"id"`
	Title        string     `json:"title"`
	Type         string     `json:"type"` // "monthly_event" or "timeline"
	Date         time.Time  `json:"date"`
	Description  string     `json:"description,omitempty"`
	Link         string     `json:"link,omitempty"`
	ProgendaId   *uuid.UUID `json:"progenda_id,omitempty"`
	ProgendaName string     `json:"progenda_name,omitempty"`
}
