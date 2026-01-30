package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// DefaultTake returns a default value for the take parameter if the provided value is less than or equal to 0.
// The default take value is 10.
func DefaultTake(i int) int {
	if i <= 0 {
		return 10
	}

	return i
}

// ToInt converts a string to an integer.
// If the conversion fails, it returns 0.
func ToInt(i string) int {
	res, err := strconv.Atoi(i)
	if err != nil {
		return 0
	}
	return res
}

func RandomData(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code[i] = charset[num.Int64()]
	}

	return string(code), nil
}
