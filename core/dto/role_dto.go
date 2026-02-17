package dto

import "github.com/google/uuid"

type CreateRoleRequest struct {
	Rank        string `json:"rank" binding:"required"`
	Index       int    `json:"index" binding:"required"`
	Description string `json:"description"`
}

type UpdateRoleRequest struct {
	Rank        string `json:"rank"`
	Index       int    `json:"index"`
	Description string `json:"description"`
}

type RoleResponse struct {
	Id          uuid.UUID `json:"id"`
	Rank        string    `json:"rank"`
	Index       int       `json:"index"`
	Description string    `json:"description"`
}
