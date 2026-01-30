package entity

import (
	"github.com/Flexoo-Academy/Golang-Template/internal/dto"
	"github.com/google/uuid"
)

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

type User struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name       string    `json:"name" gorm:"not null"`
	Email      string    `json:"email" gorm:"not null"`
	Password   string    `json:"password" gorm:"not null"`
	IsVerified bool      `json:"is_verified" gorm:"default:false;not null"`
	Role       Role      `json:"role" gorm:"default:USER;not null"`

	Timestamp
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ToInfo() dto.PersonalInfo {
	return dto.PersonalInfo{
		ID:    u.ID.String(),
		Name:  u.Name,
		Email: u.Email,
		Role:  string(u.Role),
	}
}

