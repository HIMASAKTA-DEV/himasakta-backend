package repository

import (
	"errors"
	"fmt"

	"reflect"
	"strings"

	mylog "github.com/Flexoo-Academy/Golang-Template/internal/pkg/logger"
	"github.com/Flexoo-Academy/Golang-Template/internal/pkg/meta"
	"gorm.io/gorm"
)

var (
	ErrSortBy           = errors.New("invalid sort (must be 'asc' or 'desc')")
	ErrInvalidTypeModel = errors.New("invalid type model")
	ErrInvalidField     = errors.New("invalid filter or sort field")
)

type MetaService struct {
	Filter map[string]string
	Sorter map[string]string
	DB     *gorm.DB
}

type Option func(*MetaService)

// WithFilters initializes and applies filters and sorting to a GORM query.
// It uses the provided metadata (`meta.Meta`) and optional configurations (`opts`)
// to customize the filtering and sorting behavior.
//
// Dont forget to set model in GORM query tx.Model(enitity{})
func WithFilters(db *gorm.DB, m *meta.Meta, opts ...Option) *gorm.DB {
	metaService := MetaService{
		Filter: make(map[string]string),
		Sorter: make(map[string]string),
		DB:     db,
	}

	for _, opt := range opts {
		opt(&metaService)
	}

	// for i, v := range metaService.Filter {
	// 	fmt.Println(i + " " + v)
	// }

	return metaService.buildFilter(db, m)
}

// AddModels adds entities to the service, allowing them to bypass filtering and sorting rules.
// The table prefix is used during joins to prevent ambiguous column references in the query.
func AddModels(model any) Option {
	return func(ms *MetaService) {
		stmt := &gorm.Statement{DB: ms.DB}
		if err := stmt.Parse(model); err != nil {
			mylog.Errorln(err)
		}
		tableName := stmt.Schema.Table

		v := reflect.TypeOf(model)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if v.Kind() == reflect.Struct {
			for i := 0; i < v.NumField(); i++ {
				field := v.Field(i)
				jsonTag := field.Tag.Get("json")
				if jsonTag == "" {
					jsonTag = field.Name
				}

				fullField := fmt.Sprintf("%s.%s", tableName, jsonTag)

				if field.Anonymous {
					addEmbeddedFields(field.Type, ms, tableName)
				} else {
					switch field.Type.Kind() {
					case reflect.String:
						ms.Filter[jsonTag] = fmt.Sprintf("%s ILIKE ?", fullField)
						ms.Sorter[jsonTag] = fullField
					default:
						ms.Filter[jsonTag] = fmt.Sprintf("%s = ?", fullField)
						ms.Sorter[jsonTag] = fullField
					}
				}
			}
		}
	}
}

// AddCustomField adds a custom field to the service, allowing for more flexible query building.
//
// filterQuery is the condition used in GORM queries to apply filtering.
// alias specifies the column to use when joining tables, ensuring it is included in the query.
// The table prefix is used during joins to avoid ambiguous column references in the query.
//
// Example usage:
// AddCustomField("payment_at", "transactions.updated_at = ?", "transactions.updated_at"),
func AddCustomField(field string, filterquery string, alias ...string) Option {
	return func(ms *MetaService) {
		ms.Filter[field] = filterquery
		ms.Sorter[field] = field
		if len(alias) > 0 {
			ms.Sorter[field] = alias[0]
		}
	}
}

func (ms *MetaService) buildFilter(db *gorm.DB, meta *meta.Meta) *gorm.DB {
	query := db

	filterBy := strings.Split(meta.FilterBy, ",")
	filters := strings.Split(meta.Filter, ",")

	for i, field := range filterBy {
		if i >= len(filters) {
			break
		} else if filters[i] == "" {
			continue
		}

		if condition, ok := ms.Filter[field]; ok {
			var filterValue string
			if i < len(filters) {
				filterValue = filters[i]
			}

			if strings.Contains(strings.ToUpper(condition), "ILIKE") {
				if filterValue != "" {
					filterValue = fmt.Sprintf("%%%s%%", strings.ToLower(filterValue))
				} else {
					filterValue = "%%"
				}
			}

			questionMarks := strings.Count(condition, "?")

			filterValues := make([]interface{}, questionMarks)
			for j := range filterValues {
				filterValues[j] = filterValue
			}

			query = query.Where(condition, filterValues...)
			if err := query.Error; err != nil {
				return query
			}
		} else if field != "" {
			query.Error = ErrInvalidField
			return query
		}
	}

	if meta.SortBy != "" {
		if _, ok := ms.Sorter[meta.SortBy]; !ok {
			query.Error = ErrInvalidTypeModel
			return query
		}

		if meta.Sort != "asc" && meta.Sort != "desc" {
			query.Error = ErrSortBy
			return query
		}

		order := fmt.Sprintf("%s %s", ms.Sorter[meta.SortBy], meta.Sort)

		query = query.Order(order)
		if err := query.Error; err != nil {
			return query
		}
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return query
	}

	meta.Count(int(totalCount))
	skip, limit := meta.GetSkipAndLimit()
	query = query.Scopes(paginate(skip, limit))
	if err := query.Error; err != nil {
		return query
	}
	return query
}

func paginate(page, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(page).Limit(perPage)
	}
}

func addEmbeddedFields(embedType reflect.Type, ms *MetaService, tablePrefix string) {
	if embedType.Kind() == reflect.Ptr {
		embedType = embedType.Elem()
	}
	if embedType.Kind() == reflect.Struct {
		for i := 0; i < embedType.NumField(); i++ {
			field := embedType.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = field.Name
			}

			fullField := fmt.Sprintf("%s.%s", tablePrefix, jsonTag)

			switch field.Type.Kind() {
			case reflect.String:
				ms.Filter[jsonTag] = fmt.Sprintf("%s ILIKE ?", fullField)
				ms.Sorter[jsonTag] = fullField
			default:
				ms.Filter[jsonTag] = fmt.Sprintf("%s = ?", fullField)
				ms.Sorter[jsonTag] = fullField
			}
		}
	}
}

