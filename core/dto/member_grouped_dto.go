package dto

import "github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"

type GroupedMemberResponse struct {
	Rank    string          `json:"rank"`
	Members []entity.Member `json:"members"`
}
