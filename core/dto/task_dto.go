package dto

import "time"

type (
	CreateTaskRequest struct {
		PhotoUrl    *string    `json:"photo_url" binding:""`
		Description string     `json:"description" binding:"required"`
		Deadline    *time.Time `json:"deadline" binding:"required"`
		Status      string     `json:"status" binding:"required"`
	}

	UpdateTaskRequest struct {
		PhotoUrl    *string    `json:"photo_url" binding:""`
		Description string     `json:"description" binding:"required"`
		Deadline    *time.Time `json:"deadline" binding:"required"`
		Status      string     `json:"status" binding:"required"`
	}
)
