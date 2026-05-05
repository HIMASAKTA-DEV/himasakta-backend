package dbimport

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

func ImportDB(db *gorm.DB, zipPath string, storageTarget string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	var sqlFile *zip.File
	var storageFiles []*zip.File
	var metaFile *zip.File

	for _, f := range r.File {
		if f.Name == "metadata.json" {
			metaFile = f
		} else if f.Name == "database.sql" {
			sqlFile = f
		} else if strings.HasPrefix(f.Name, "storage/") && !f.FileInfo().IsDir() {
			storageFiles = append(storageFiles, f)
		}
	}

	if sqlFile == nil {
		return fmt.Errorf("database.sql not found in zip archive")
	}

	log.Println("Importing database records...")
	if err := importSQL(db, sqlFile); err != nil {
		return fmt.Errorf("failed to import sql: %w", err)
	}

	log.Printf("Importing %d storage files to %s...", len(storageFiles), storageTarget)
	if storageTarget == "s3" {
		if err := importS3Storage(storageFiles); err != nil {
			return err
		}
	} else if storageTarget == "local" {
		if err := importLocalStorage(storageFiles); err != nil {
			return err
		}
	}

	log.Println("Import complete.")
	return nil
}

func importSQL(db *gorm.DB, sqlFile *zip.File) error {
	rc, err := sqlFile.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	sqlBytes, err := io.ReadAll(rc)
	if err != nil {
		return fmt.Errorf("failed to read database.sql: %w", err)
	}

	sqlString := string(sqlBytes)

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying db: %w", err)
	}

	_, err = sqlDB.Exec(sqlString)
	if err != nil {
		return fmt.Errorf("failed to execute sql dump: %w", err)
	}

	return nil
}

func importLocalStorage(files []*zip.File) error {
	uploadBasePath := "assets/uploads"
	for _, f := range files {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(f.Name, "storage/")
		targetPath := filepath.Join(uploadBasePath, filepath.FromSlash(relPath))

		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			rc.Close()
			return err
		}

		dst, err := os.Create(targetPath)
		if err != nil {
			rc.Close()
			return err
		}

		if _, err := io.Copy(dst, rc); err != nil {
			dst.Close()
			rc.Close()
			return err
		}

		dst.Close()
		rc.Close()
	}
	return nil
}

func importS3Storage(files []*zip.File) error {
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
			accessKey, secretKey, "",
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

	for _, f := range files {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return err
		}

		key := strings.TrimPrefix(f.Name, "storage/")

		var sample []byte
		if len(content) > 512 {
			sample = content[:512]
		} else {
			sample = content
		}
		mimetype := http.DetectContentType(sample)

		if strings.HasSuffix(key, ".svg") {
			mimetype = "image/svg+xml"
		} else if strings.HasSuffix(key, ".json") {
			mimetype = "application/json"
		}

		_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(key),
			Body:        bytes.NewReader(content),
			ContentType: aws.String(mimetype),
		})

		if err != nil {
			log.Printf("Failed to upload %s to S3: %v", key, err)
		}
	}

	return nil
}
