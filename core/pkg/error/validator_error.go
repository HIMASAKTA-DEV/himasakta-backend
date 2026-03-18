package myerror

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseValidationError(err error) string {
	if err == nil {
		return ""
	}

	if errs, ok := err.(validator.ValidationErrors); ok {
		var errMsgs []string
		for _, e := range errs {
			field := e.Field()
			tag := e.Tag()

			switch tag {
			case "required":
				errMsgs = append(errMsgs, fmt.Sprintf("Field '%s' belum diisi", field))
			case "email":
				errMsgs = append(errMsgs, fmt.Sprintf("Field '%s' harus berupa email yang valid", field))
			case "min":
				errMsgs = append(errMsgs, fmt.Sprintf("Field '%s' minimal %s karakter", field, e.Param()))
			case "max":
				errMsgs = append(errMsgs, fmt.Sprintf("Field '%s' maksimal %s karakter", field, e.Param()))
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("Field '%s' tidak valid (%s)", field, tag))
			}
		}
		return strings.Join(errMsgs, ", ")
	}
	return err.Error()
}
