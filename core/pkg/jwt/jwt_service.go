package myjwt

import (
	"time"
)

type JWT interface {
	CreateToken(userId string, username string) (string, error)
	ValidateToken(token string) (bool, error)
	GetClaims(token string) (map[string]string, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWT() JWT {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    getIssuer(),
	}
}

func (j *jwtService) CreateToken(userId string, username string) (string, error) {
	payload := map[string]string{
		"user_id":  userId,
		"username": username,
	}
	return GenerateToken(payload, 24*time.Hour)
}

func (j *jwtService) ValidateToken(token string) (bool, error) {
	return IsValid(token)
}

func (j *jwtService) GetClaims(token string) (map[string]string, error) {
	return GetPayloadInsideToken(token)
}
