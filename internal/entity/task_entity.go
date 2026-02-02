package entity

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "PENDING"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted  TaskStatus = "COMPLETED"
)

type Task struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PhotoUrl    *string    `json:"photo_url" gorm:""`
	Description string     `json:"description" gorm:"not null"`
	Status      TaskStatus `json:"status" gorm:"not null"`
	Deadline    *time.Time `json:"deadline" gorm:"not null"`

	Timestamp
}

func (t *Task) TableName() string {
	return "tasks"
}
