package dto

type (
	CreateUserRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required,oneof=USER ADMIN"`
	}

	UpdateUserRequest struct {
		ID       string  `json:"id"`
		Name     string  `json:"name" binding:"required"`
		Email    string  `json:"email" binding:"required,email"`
		Password *string `json:"password" binding:""`
		Role     string  `json:"role" binding:"required,oneof=USER ADMIN"`
	}

	UserResponse struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
		Role       string `json:"role"`
	}
)
