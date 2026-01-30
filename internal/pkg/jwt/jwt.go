package myjwt

import (
	"errors"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload map[string]string, ExpiredAt time.Duration) (string, error) {
	expiredAt := time.Now().Add(ExpiredAt).Unix()

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["iss"] = getIssuer()

	for i, v := range payload {
		claims[i] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func GetPayloadInsideToken(tokenString string) (map[string]string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(getSecretKey()), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	payload := make(map[string]string)
	for key, value := range claims {
		if strValue, ok := value.(string); ok {
			payload[key] = strValue
		}
	}

	return payload, nil
}

func IsValid(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(getSecretKey()), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return false, errors.New("token expired")
		}
		return false, err
	}

	if !token.Valid {
		return false, jwt.ErrInvalidKey
	}

	return true, nil
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "Template"
	}
	return secretKey
}

func getIssuer() string {
	return "CRS Backend 2025"
}

