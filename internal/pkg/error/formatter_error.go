package myerror

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
)

func GetErrBodyRequest(actErr error, v any) error {
	if validationErrors, ok := actErr.(validator.ValidationErrors); ok {
		errBody := New(Descriptive(validationErrors, v), http.StatusBadRequest)
		return Wrap(errBody, ErrBodyRequest)
	}
	actErr = New(actErr.Error(), http.StatusBadRequest)
	return Wrap(actErr, ErrBodyRequest)
}

func Wrap(actErr error, prodErr error) error {
	mode := os.Getenv("APP_MODE")
	switch mode {
	case "dev":
		return actErr
	default:
		return prodErr
	}
}
