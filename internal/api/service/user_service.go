package service

import (
	"github.com/Flexoo-Academy/Golang-Template/internal/api/repository"
	"gorm.io/gorm"
)

type (
	UserService interface {
	}

	userService struct {
		userRepository repository.UserRepository

		db *gorm.DB
	}
)

func NewUser(userRepository repository.UserRepository,
	db *gorm.DB) UserService {
	return &userService{
		userRepository: userRepository,
		db:             db,
	}
}

