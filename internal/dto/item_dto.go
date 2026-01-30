package dto

type (
	CreateItemRequest struct {
		CategoryCode string `json:"category_code" binding:"required"`
		CategoryName string `json:"category_name" binding:"required"`
		Group        string `json:"group" binding:"required"`
		AccountCode  string `json:"account_code" binding:"required"`
		Description  string `json:"description" binding:"required"`
	}

	UpdateItemRequest struct {
		ID           string `json:"-"`
		CategoryCode string `json:"category_code"`
		CategoryName string `json:"category_name"`
		Group        string `json:"group"`
		AccountCode  string `json:"account_code"`
		Description  string `json:"description"`
	}
)
