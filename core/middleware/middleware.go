package middleware

import (
	myjwt "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/jwt"
	"gorm.io/gorm"
)

type Middleware struct {
	db         *gorm.DB
	jwtService myjwt.JWT
}

func New(db *gorm.DB) Middleware {
	return Middleware{
		db:         db,
		jwtService: myjwt.NewJWT(),
	}
}
