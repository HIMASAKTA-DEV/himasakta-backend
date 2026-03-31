package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"sync"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/utils"
)

const uploadBasePath = "assets/uploads"

type localStorage struct {
	appURL     string
	actions    []action
	isRollback bool
}

func NewLocalStorage() (FileStorage, error) {
	appURL := os.Getenv("APP_URL")
	if appURL == "" {
		appURL = "http://localhost:8080"
	}
	appURL = strings.TrimRight(appURL, "/")

	if err := os.MkdirAll(uploadBasePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	return &localStorage{
		appURL:     appURL,
		actions:    nil,
		isRollback: false,
	}, nil
}

func (l *localStorage) UploadFile(filename string, f *multipart.FileHeader, folderName string, mv ...string) (string, error) {
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
		allowed := false
		for _, m := range mv {
			if mimetype == m {
				allowed = true
				break
			}
		}
		if !allowed {
			return "", ErrInvalidTypeFile
		}
	}

	objectKey := fmt.Sprintf("%s/%s", folderName, filename)
	dirPath := fmt.Sprintf("%s/%s", uploadBasePath, folderName)

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory %s: %w", dirPath, err)
	}

	filePath := fmt.Sprintf("%s/%s", uploadBasePath, objectKey)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		os.Remove(filePath)
		return "", err
	}

	l.actions = append(l.actions, action{actionType: "upload", key: objectKey})
	return objectKey, nil
}

func (l *localStorage) UpdateFile(objectKey string, f *multipart.FileHeader, mv ...string) (string, error) {
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
		allowed := false
		for _, m := range mv {
			if mimetype == m {
				allowed = true
				break
			}
		}
		if !allowed {
			return "", ErrInvalidTypeFile
		}
	}

	filePath := fmt.Sprintf("%s/%s", uploadBasePath, objectKey)

	parts := strings.Split(objectKey, "/")
	if len(parts) > 1 {
		dirPath := fmt.Sprintf("%s/%s", uploadBasePath, strings.Join(parts[:len(parts)-1], "/"))
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return "", err
		}
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}

	return objectKey, nil
}

func (l *localStorage) DeleteFile(objectKey string) error {
	filePath := fmt.Sprintf("%s/%s", uploadBasePath, objectKey)
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func (l *localStorage) GetPublicLink(objectKey string) string {
	return fmt.Sprintf("%s/api/static/%s", l.appURL, objectKey)
}

func (l *localStorage) GetObjectKeyFromLink(link string) string {
	prefix := fmt.Sprintf("%s/api/static/", l.appURL)
	if !strings.HasPrefix(link, prefix) {
		return ""
	}
	return strings.TrimPrefix(link, prefix)
}

func (l *localStorage) Begin() FileStorage {
	l.actions = []action{}
	l.isRollback = true
	return l
}

func (l *localStorage) Commit() {
	l.actions = nil
	l.isRollback = false
}

func (l *localStorage) Rollback() {
	var wg sync.WaitGroup
	errCh := make(chan error, len(l.actions))

	for i := len(l.actions) - 1; i >= 0; i-- {
		act := l.actions[i]
		if act.actionType == "upload" {
			wg.Add(1)
			go func(key string) {
				defer wg.Done()
				if err := l.DeleteFile(key); err != nil {
					errCh <- fmt.Errorf("rollback failed to delete %s: %v", key, err)
				}
			}(act.key)
		}
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		panic("failed rollback local storage")
	}

	l.Commit()
}
