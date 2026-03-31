package storage

import (
	"fmt"
	"mime/multipart"
	"os"
)

type FileStorage interface {
	UploadFile(filename string, file *multipart.FileHeader, folderName string, mv ...string) (string, error)
	UpdateFile(objectKey string, f *multipart.FileHeader, mv ...string) (string, error)
	DeleteFile(objectKey string) error
	GetPublicLink(objectKey string) string
	GetObjectKeyFromLink(link string) string
	Begin() FileStorage
	Commit()
	Rollback()
}

func NewFileStorage() (FileStorage, error) {
	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "" {
		storageType = "local"
	}

	switch storageType {
	case "s3":
		return NewAwsS3()
	case "local":
		return NewLocalStorage()
	default:
		return nil, fmt.Errorf("unknown STORAGE_TYPE: %s (use 'local' or 's3')", storageType)
	}
}
