package migrations

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println(mylog.ColorizeInfo("\n=========== Start Migrate ==========="))

	mylog.Infof("Migrating Tables...")

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}

	allEntities := []interface{}{
		&entity.Role{},
		&entity.Department{},
		&entity.CabinetInfo{},
		&entity.GlobalSetting{},
		&entity.NrpWhitelist{},
		&entity.Visitor{},
		&entity.MonthlyEvent{},
		&entity.Tag{},
		&entity.Gallery{},
		&entity.Progenda{},
		&entity.Timeline{},
		&entity.Member{},
		&entity.News{},
		&entity.NewsTag{},
	}

	db.Config.DisableForeignKeyConstraintWhenMigrating = true
	if err := db.AutoMigrate(allEntities...); err != nil {
		return err
	}

	db.Config.DisableForeignKeyConstraintWhenMigrating = false
	if err := db.AutoMigrate(allEntities...); err != nil {
		return err
	}

	if err := seedAdmin(db); err != nil {
		mylog.Errorf("Failed to seed admin: %v", err)
	}

	if err := migrateImageURLs(db); err != nil {
		mylog.Errorf("Failed to migrate image URLs: %v", err)
	}

	mylog.Infof("Migration completed successfully")

	return nil
}

func seedAdmin(db *gorm.DB) error {
	var existing entity.GlobalSetting
	result := db.Where("key = ?", "auth").First(&existing)
	if result.Error == nil {
		mylog.Infof("Admin credentials already exist, skipping seed")
		return nil
	}

	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	if username == "" || password == "" {
		mylog.Infof("ADMIN_USERNAME or ADMIN_PASSWORD not set, skipping admin seed")
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	authJSON, _ := json.Marshal(map[string]string{
		"username": username,
		"password": string(hash),
	})

	setting := entity.GlobalSetting{
		Key:   "auth",
		Value: string(authJSON),
	}

	if err := db.Create(&setting).Error; err != nil {
		return err
	}

	mylog.Infof("Initial admin credentials seeded (user: %s)", username)
	return nil
}

func migrateImageURLs(db *gorm.DB) error {
	mylog.Infof("Checking and migrating image URLs to relative paths...")

	appURL := os.Getenv("APP_URL")
	if appURL == "" {
		appURL = "http://localhost:8080"
	}
	appURL = strings.TrimRight(appURL, "/")
	localPrefix := fmt.Sprintf("%s/api/static/", appURL)

	s3Prefix := ""
	if customPref := os.Getenv("S3_PUBLIC_URL_PREFIX"); customPref != "" {
		s3Prefix = customPref
	} else if bucket := os.Getenv("S3_BUCKET"); bucket != "" {
		region := os.Getenv("AWS_REGION")
		s3Prefix = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", bucket, region)
	}

	prefixes := []string{localPrefix}
	if s3Prefix != "" {
		prefixes = append(prefixes, s3Prefix)
	}

	for _, prefix := range prefixes {
		// Update galleries
		res := db.Exec(`UPDATE galleries SET image_url = '/api/static/' || SUBSTRING(image_url FROM LENGTH(?) + 1) WHERE image_url LIKE ?`, prefix, prefix+"%")
		if res.Error != nil {
			return fmt.Errorf("failed to update galleries: %w", res.Error)
		}
		if res.RowsAffected > 0 {
			mylog.Infof("Updated %d gallery rows for prefix %s", res.RowsAffected, prefix)
		}

		// Update global_settings
		var settings entity.GlobalSetting
		if err := db.Where("key = ?", "web_settings").First(&settings).Error; err == nil {
			var webSettings map[string]interface{}
			if err := json.Unmarshal([]byte(settings.Value), &webSettings); err == nil {
				changed := false
				for key, val := range webSettings {
					if str, ok := val.(string); ok {
						if strings.HasPrefix(str, prefix) {
							webSettings[key] = "/api/static/" + strings.TrimPrefix(str, prefix)
							changed = true
						}
					}
				}
				if changed {
					newJSON, _ := json.Marshal(webSettings)
					settings.Value = string(newJSON)
					db.Save(&settings)
					mylog.Infof("Updated global_settings web_settings for prefix %s", prefix)
				}
			}
		}
	}

	return nil
}

