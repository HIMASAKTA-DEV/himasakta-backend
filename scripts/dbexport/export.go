package dbexport

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	if err := os.MkdirAll("vps", 0755); err != nil {
		return fmt.Errorf("failed to create vps directory: %w", err)
	}

	timestamp := time.Now().Format("20060102_150405")
	zipFilename := fmt.Sprintf("vps/export_%s.zip", timestamp)

	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 1. Write SQL dump
	log.Printf("Exporting database to SQL dump...")
	sqlHeader, err := zipWriter.Create("database.sql")
	if err != nil {
		return err
	}
	if err := writeSQLDump(db, sqlHeader); err != nil {
		return err
	}

	// 2. Export storage
	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "" {
		storageType = "local"
	}

	log.Printf("Exporting storage files from %s...", storageType)
	if storageType == "s3" {
		if err := exportS3Storage(zipWriter); err != nil {
			return err
		}
	} else if storageType == "local" {
		if err := exportLocalStorage(zipWriter); err != nil {
			return err
		}
	} else {
		log.Printf("Unknown storage type: %s, skipping storage export", storageType)
	}

	log.Printf("Export complete: %s", zipFilename)
	return nil
}

func writeSQLDump(db *gorm.DB, w io.Writer) error {
	fmt.Fprintln(w, "-- Database export generated at", time.Now().Format(time.RFC3339))
	fmt.Fprintln(w, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	fmt.Fprintln(w, "")

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
			continue
		}

		totalRows += count
		fmt.Fprintf(w, "-- Table: %s (%d rows)\n", t.name, count)

		for i := 0; i < count; i++ {
			row := sliceVal.Index(i)
			cols, vals := extractColumnsAndValues(row)
			if len(cols) == 0 {
				continue
			}
			fmt.Fprintf(w, "INSERT INTO %s (%s) VALUES (%s) ON CONFLICT DO NOTHING;\n",
				t.name,
				strings.Join(cols, ", "),
				strings.Join(vals, ", "),
			)
		}
		fmt.Fprintln(w, "")
	}

	log.Printf("SQL Export complete (%d total rows)", totalRows)
	return nil
}

func exportLocalStorage(zw *zip.Writer) error {
	uploadBasePath := "assets/uploads"
	if _, err := os.Stat(uploadBasePath); os.IsNotExist(err) {
		log.Printf("No local uploads found at %s", uploadBasePath)
		return nil
	}

	count := 0
	err := filepath.Walk(uploadBasePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(uploadBasePath, path)
		if err != nil {
			return err
		}

		zipEntryPath := "storage/" + filepath.ToSlash(relPath)
		fileHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		fileHeader.Name = zipEntryPath
		fileHeader.Method = zip.Deflate

		writer, err := zw.CreateHeader(fileHeader)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := io.Copy(writer, file); err != nil {
			return err
		}
		count++
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to export local storage: %w", err)
	}
	log.Printf("Exported %d files from local storage", count)
	return nil
}

func exportS3Storage(zw *zip.Writer) error {
	bucket := os.Getenv("S3_BUCKET")
	region := os.Getenv("AWS_REGION")
	endpoint := os.Getenv("S3_ENDPOINT")

	accessKey := os.Getenv("AWS_ACCESS_KEY")
	if accessKey == "" {
		accessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	}
	secretKey := os.Getenv("AWS_SECRET_KEY")
	if secretKey == "" {
		secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	}

	options := []func(*config.LoadOptions) error{
		config.WithRegion(region),
	}
	if accessKey != "" && secretKey != "" {
		options = append(options, config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			"",
		)))
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), options...)
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		if endpoint != "" {
			o.BaseEndpoint = aws.String(endpoint)
		}
	})

	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})

	count := 0
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return fmt.Errorf("failed to list objects: %w", err)
		}

		for _, obj := range page.Contents {
			key := *obj.Key

			resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(key),
			})
			if err != nil {
				log.Printf("Failed to download %s: %v", key, err)
				continue
			}

			zipEntryPath := "storage/" + key
			writer, err := zw.Create(zipEntryPath)
			if err != nil {
				resp.Body.Close()
				return err
			}

			if _, err := io.Copy(writer, resp.Body); err != nil {
				resp.Body.Close()
				return err
			}
			resp.Body.Close()
			count++
		}
	}

	log.Printf("Exported %d files from S3", count)
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
