package dto

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
)

// ... existing DTOs ...

type MemberGroupResponse struct {
	Role    entity.Role     `json:"role"`
	Members []entity.Member `json:"members"`
}
