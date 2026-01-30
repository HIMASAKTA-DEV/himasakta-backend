package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const PATH = "public/uploads"

func UploadFile(file *multipart.FileHeader, path string) error {
	parts := strings.Split(path, "/")
	fileID := parts[len(parts)-1]
	dirPath := PATH
	fullDirPath := fmt.Sprintf("%s/%s", dirPath, strings.Join(parts[:len(parts)-1], "/"))

	if _, err := os.Stat(fullDirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(fullDirPath, 0777); err != nil {
			return err
		}
	}

	filePath := fmt.Sprintf("%s/%s", fullDirPath, fileID)

	uploadedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadedFile.Close()

	targetFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, uploadedFile)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFile(path string) error {
	if err := os.Remove(fmt.Sprintf("%s/%s", PATH, path)); err != nil {
		return err
	}

	return nil
}

func GetExtensions(filename string) string {
	ext := strings.Split(filename, ".")
	return ext[len(ext)-1]
}

func GetMimetype(f multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	mimeType := http.DetectContentType(buffer)

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	return mimeType, nil
}
