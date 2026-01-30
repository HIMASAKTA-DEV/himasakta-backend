package form

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseForm(c *gin.Context, obj any) error {
	v := reflect.TypeOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return errors.New("invalid type form model")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fType := field.Type.Kind()
		formTag := field.Tag.Get("form")

		if fType != reflect.Slice && fType != reflect.Struct {
			fieldValue := reflect.ValueOf(obj).Elem().Field(i)
			setter(c, fType, fieldValue, formTag)
		} else {
			switch fType {
			case reflect.Slice:
				v, err := parseFormArray(c, reflect.New(field.Type).Interface(), formTag)
				if err != nil {
					return err
				}
				reflect.ValueOf(obj).Elem().Field(i).Set(v)
			case reflect.Struct:
				ParseForm(c, reflect.New(field.Type).Interface())
			}
		}
	}

	return nil
}

func parseFormArray(c *gin.Context, obj any, tag string) (reflect.Value, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice {
		return reflect.Value{}, errors.New("invalid type form model")
	}

	total := 0
	indexes := make(map[int]bool)
	for key := range c.Request.PostForm {
		if !strings.HasPrefix(key, tag) {
			continue
		}
		match := regexp.MustCompile(`\[(.*?)\]`).FindString(key)
		if match == "" {
			continue
		}
		idx, _ := strconv.Atoi(strings.Trim(match, "[]"))
		if _, ok := indexes[idx]; !ok {
			indexes[idx] = true
			total++
		}
	}

	slices := reflect.MakeSlice(v.Type(), total, total)
	for i := 0; i < total; i++ {
		sv := slices.Index(i)
		for j := 0; j < sv.NumField(); j++ {
			field := sv.Field(j)
			fType := field.Type().Kind()
			formTag := sv.Type().Field(j).Tag.Get("form")
			key := fmt.Sprintf("%s[%d][%s]", tag, i, formTag)
			setter(c, fType, field, key)
		}
	}

	return slices, nil
}

func setter(c *gin.Context, t reflect.Kind, v reflect.Value, key string) {
	formValue := c.PostForm(key)
	if v.Type().String() == "*multipart.FileHeader" {
		file, _ := c.FormFile(key)
		v.Set(reflect.ValueOf(file))
		return
	}
	switch t {
	case reflect.String:
		v.SetString(formValue)
	case reflect.Int:
		intValue, _ := strconv.Atoi(formValue)
		v.SetInt(int64(intValue))
	}
}
