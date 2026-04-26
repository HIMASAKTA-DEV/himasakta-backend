package dbexport

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"gorm.io/gorm"
)

type tableExport struct {
	name  string
	model interface{}
	slice interface{}
}

func getAllTables() []tableExport {
	return []tableExport{
		{"global_settings", entity.GlobalSetting{}, &[]entity.GlobalSetting{}},
		{"roles", entity.Role{}, &[]entity.Role{}},
		{"departments", entity.Department{}, &[]entity.Department{}},
		{"cabinet_infos", entity.CabinetInfo{}, &[]entity.CabinetInfo{}},
		{"galleries", entity.Gallery{}, &[]entity.Gallery{}},
		{"progendas", entity.Progenda{}, &[]entity.Progenda{}},
		{"timelines", entity.Timeline{}, &[]entity.Timeline{}},
		{"members", entity.Member{}, &[]entity.Member{}},
		{"monthly_events", entity.MonthlyEvent{}, &[]entity.MonthlyEvent{}},
		{"tags", entity.Tag{}, &[]entity.Tag{}},
		{"news", entity.News{}, &[]entity.News{}},
		{"news_tags", entity.NewsTag{}, &[]entity.NewsTag{}},
		{"nrp_whitelists", entity.NrpWhitelist{}, &[]entity.NrpWhitelist{}},
		{"visitors", entity.Visitor{}, &[]entity.Visitor{}},
	}
}

func ExportDB(db *gorm.DB) error {
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("export_%s.sql", timestamp)
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	fmt.Fprintln(f, "-- Database export generated at", time.Now().Format(time.RFC3339))
	fmt.Fprintln(f, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	fmt.Fprintln(f, "")

	tables := getAllTables()
	totalRows := 0

	for _, t := range tables {
		slicePtr := t.slice
		result := db.Table(t.name).Find(slicePtr)
		if result.Error != nil {
			log.Printf("Warning: failed to read %s: %v", t.name, result.Error)
			continue
		}

		sliceVal := reflect.ValueOf(slicePtr).Elem()
		count := sliceVal.Len()
		if count == 0 {
			log.Printf("  %s: 0 rows (skipped)", t.name)
			continue
		}

		log.Printf("  %s: %d rows", t.name, count)
		totalRows += count

		fmt.Fprintf(f, "-- Table: %s (%d rows)\n", t.name, count)

		for i := 0; i < count; i++ {
			row := sliceVal.Index(i)
			cols, vals := extractColumnsAndValues(row)
			if len(cols) == 0 {
				continue
			}
			fmt.Fprintf(f, "INSERT INTO %s (%s) VALUES (%s) ON CONFLICT DO NOTHING;\n",
				t.name,
				strings.Join(cols, ", "),
				strings.Join(vals, ", "),
			)
		}
		fmt.Fprintln(f, "")
	}

	log.Printf("Export complete: %s (%d total rows)", filename, totalRows)
	return nil
}

func extractColumnsAndValues(row reflect.Value) ([]string, []string) {
	var cols []string
	var vals []string

	rowType := row.Type()
	for i := 0; i < row.NumField(); i++ {
		field := rowType.Field(i)
		val := row.Field(i)

		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			for j := 0; j < val.NumField(); j++ {
				subField := field.Type.Field(j)
				subVal := val.Field(j)
				col := getColumnName(subField)
				if col == "" || col == "-" {
					continue
				}
				cols = append(cols, col)
				vals = append(vals, formatValue(subVal))
			}
			continue
		}

		col := getColumnName(field)
		if col == "" || col == "-" {
			continue
		}

		if field.Type.Kind() == reflect.Slice {
			continue
		}
		if field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct {
			elemType := field.Type.Elem()
			typeName := elemType.Name()
			if typeName != "UUID" && typeName != "Time" {
				continue
			}
		}

		cols = append(cols, col)
		vals = append(vals, formatValue(val))
	}
	return cols, vals
}

func getColumnName(field reflect.StructField) string {
	gormTag := field.Tag.Get("gorm")
	if gormTag != "" {
		for _, part := range strings.Split(gormTag, ";") {
			part = strings.TrimSpace(part)
			if strings.HasPrefix(part, "column:") {
				return strings.TrimPrefix(part, "column:")
			}
		}
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag != "" && jsonTag != "-" {
		parts := strings.Split(jsonTag, ",")
		return parts[0]
	}

	return ""
}

func formatValue(val reflect.Value) string {
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "NULL"
		}
		val = val.Elem()
	}

	switch v := val.Interface().(type) {
	case time.Time:
		if v.IsZero() {
			return "NULL"
		}
		return fmt.Sprintf("'%s'", v.Format("2006-01-02 15:04:05"))
	case string:
		escaped := strings.ReplaceAll(v, "'", "''")
		return fmt.Sprintf("'%s'", escaped)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%v", v)
	case bool:
		if v {
			return "TRUE"
		}
		return "FALSE"
	default:
		s := fmt.Sprintf("%v", v)
		escaped := strings.ReplaceAll(s, "'", "''")
		return fmt.Sprintf("'%s'", escaped)
	}
}
