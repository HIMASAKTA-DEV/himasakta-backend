package myjwt

type JWT interface {
	CreateToken(userId string, username string, role string) (string, error)
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

func (j *jwtService) CreateToken(userId string, username string, role string) (string, error) {
	payload := map[string]string{
		"user_id":  userId,
		"username": username,
		"role":     role,
	}
	return GenerateToken(payload, GetExpirationDuration())
}

func (j *jwtService) ValidateToken(token string) (bool, error) {
	return IsValid(token)
}

func (j *jwtService) GetClaims(token string) (map[string]string, error) {
	return GetPayloadInsideToken(token)
}
