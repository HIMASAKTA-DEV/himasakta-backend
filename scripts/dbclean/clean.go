package dbclean

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gorm.io/gorm"
)

func CleanData(db *gorm.DB) error {
	log.Println("WARNING: Executing complete database wipe...")
	
	// Drop schema and recreate
	err := db.Exec("DROP SCHEMA public CASCADE;").Error
	if err != nil {
		return fmt.Errorf("failed to drop schema: %w", err)
	}

	err = db.Exec("CREATE SCHEMA public;").Error
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	err = db.Exec("GRANT ALL ON SCHEMA public TO postgres;").Error
	if err != nil {
		log.Printf("Warning: failed to grant permissions: %v", err)
	}

	err = db.Exec("GRANT ALL ON SCHEMA public TO public;").Error
	if err != nil {
		log.Printf("Warning: failed to grant public permissions: %v", err)
	}

	log.Println("Database schema wiped successfully.")

	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "" {
		storageType = "local"
	}

	log.Printf("Cleaning %s storage...", storageType)
	if storageType == "local" {
		uploadBasePath := "assets/uploads"
		os.RemoveAll(uploadBasePath)
		os.MkdirAll(uploadBasePath, 0755)
		log.Println("Local storage wiped successfully.")
	} else if storageType == "s3" {
		if err := cleanS3(); err != nil {
			return fmt.Errorf("failed to clean s3: %w", err)
		}
	} else {
		log.Printf("Unknown storage type %s. Skipping storage wipe.", storageType)
	}

	return nil
}

func cleanS3() error {
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
		return err
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
			return err
		}

		for _, obj := range page.Contents {
			_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
				Bucket: aws.String(bucket),
				Key:    obj.Key,
			})
			if err != nil {
				log.Printf("Failed to delete %s: %v", *obj.Key, err)
				continue
			}
			count++
		}
	}

	log.Printf("Wiped %d objects from S3 successfully.", count)
	return nil
}
