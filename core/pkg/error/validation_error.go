package myerror

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var tagDescriptions = map[string]string{
	"required": "This field is required.",
	"email":    "Must be a valid email address.",
	"gte":      "The value must be greater than or equal to %s.",
	"lte":      "The value must be less than or equal to %s.",
	"len":      "The length of the string must be exactly %s.",
	"http_url": "Must be a valid url format.",
}

func Descriptive(verr validator.ValidationErrors, obj any) string {
	errorMsg := []string{}
	valType := reflect.TypeOf(obj)
	for _, f := range verr {
		fieldName := f.Field()
		field, ok := valType.FieldByName(fieldName)
		if !ok {
			continue
		}

		fieldJSONName, _ := field.Tag.Lookup("json")
		description := tagDescriptions[f.ActualTag()]
		if description == "" {
			description = f.ActualTag()
		}

		if f.Param() != "" {
			description = fmt.Sprintf(description, f.Param())
		}

		errorMsg = append(errorMsg, fmt.Sprintf("field '%s': %s", fieldJSONName, description))
	}

	return strings.Join(errorMsg, " ")
}
