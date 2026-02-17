package dto

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Level       int    `json:"level" binding:"required"`
	Description string `json:"description"`
}

type UpdateRoleRequest struct {
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}
