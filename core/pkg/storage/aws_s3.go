package storage

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"sync"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type (
	AwsS3 interface {
		UploadFile(filename string, file *multipart.FileHeader, folderName string, mv ...string) (string, error)
		UpdateFile(objectKey string, f *multipart.FileHeader, mv ...string) (string, error)
		DeleteFile(objectKey string) error
		GetPublicLink(objectKey string) string
		GetObjectKeyFromLink(link string) string
		Begin() AwsS3
		Commit()
		Rollback()
	}

	action struct {
		actionType string
		key        string
	}

	awsS3 struct {
		client     *s3.Client
		bucket     string
		region     string
		endpoint   string
		actions    []action
		isRollback bool
	}
)

var (
	ErrInvalidTypeFile = errors.New("invalid type file")
)

func NewAwsS3() AwsS3 {
	bucket := os.Getenv("S3_BUCKET")
	region := os.Getenv("AWS_REGION")
	endpoint := os.Getenv("S3_ENDPOINT")

	var (
		cfg aws.Config
		err error
	)

	// Try custom env vars first, then fallback to standard AWS SDK env vars
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
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			"",
		)),
	}

	cfg, err = config.LoadDefaultConfig(context.TODO(), options...)
	if err != nil {
		panic(fmt.Sprintf("failed to load AWS configuration: %v", err))
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		if endpoint != "" {
			o.BaseEndpoint = aws.String(endpoint)
		}
	})

	return &awsS3{
		client:     client,
		bucket:     bucket,
		region:     region,
		endpoint:   endpoint,
		actions:    nil,
		isRollback: false,
	}
}

// Gunakan untuk mengupload file ke s3 dimana defaultnya mengizinkan semua jenis mimetypa
func (a *awsS3) UploadFile(filename string, f *multipart.FileHeader, folderName string, mv ...string) (string, error) {
	file, err := f.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	mimetype, err := utils.GetMimetype(file)
	if err != nil {
		return "", err
	}

	if len(mv) > 0 {
		flag := false
		for _, m := range mv {
			if mimetype == m {
				flag = true
				break
			}
		}

		if !flag {
			return "", ErrInvalidTypeFile
		}
	}

	objectKey := fmt.Sprintf("%s/%s", folderName, filename)

	_, err = a.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(a.bucket),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(mimetype),
	})
	if err != nil {
		return "", err
	}

	a.actions = append(a.actions, action{actionType: "upload", key: objectKey})

	return objectKey, nil
}

// Gunakan untuk mengupdate file ke s3 dimana defaultnya mengizinkan semua jenis mimetypa
func (a *awsS3) UpdateFile(objectKey string, f *multipart.FileHeader, mv ...string) (string, error) {
	file, err := f.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	mimetype, err := utils.GetMimetype(file)
	if err != nil {
		return "", err
	}

	if len(mv) > 0 {
		flag := false
		for _, m := range mv {
			if mimetype == m {
				flag = true
				break
			}
		}

		if !flag {
			return "", ErrInvalidTypeFile
		}
	}

	_, err = a.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(a.bucket),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(mimetype),
	})
	if err != nil {
		return "", err
	}

	return objectKey, nil
}

// Gunakan untuk menghapus file ke s3 dimana defaultnya mengizinkan semua jenis mimetypa
func (a *awsS3) DeleteFile(objectKey string) error {
	_, err := a.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *awsS3) GetObjectKeyFromLink(link string) string {
	pref := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", a.bucket, a.region)
	if a.endpoint != "" {
		// If using Supabase, public links usually follow this pattern or similar
		// We allow S3_PUBLIC_URL_PREFIX for flexibility
		customPref := os.Getenv("S3_PUBLIC_URL_PREFIX")
		if customPref != "" {
			pref = customPref
		}
	}

	if !strings.HasPrefix(link, pref) {
		return ""
	}

	objectKey := strings.TrimPrefix(link, pref)
	return objectKey
}

func (a *awsS3) GetPublicLink(objectKey string) string {
	// Default AWS S3 format
	publicURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", a.bucket, a.region, objectKey)

	// If a custom public URL prefix is provided
	if customPref := os.Getenv("S3_PUBLIC_URL_PREFIX"); customPref != "" {
		return fmt.Sprintf("%s%s", customPref, objectKey)
	}

	return publicURL
}

func (a *awsS3) Begin() AwsS3 {
	a.actions = []action{}
	a.isRollback = true

	return a
}

func (a *awsS3) Commit() {
	a.actions = nil
	a.isRollback = false
}

func (a *awsS3) Rollback() {
	var wg sync.WaitGroup
	errCh := make(chan error, len(a.actions))

	for i := len(a.actions) - 1; i >= 0; i-- {
		action := a.actions[i]
		switch action.actionType {
		case "upload":
			wg.Add(1)
			go func(key string) {
				defer wg.Done()
				if err := a.DeleteFile(key); err != nil {
					errCh <- fmt.Errorf("failed to delete file %s: %v", key, err)
				}
			}(action.key)
		}
	}

	wg.Wait()

	close(errCh)

	var errors []error
	for err := range errCh {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		panic("failed rollback")
	}

	a.Commit()
}
