package db

import (
	"fmt"

	"os"

	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	var DBDSN string

	// Vercel / Supabase integration usually provides POSTGRES_URL
	if postgresURL := os.Getenv("POSTGRES_URL"); postgresURL != "" {
		DBDSN = postgresURL
	} else if pgHost := os.Getenv("POSTGRES_HOST"); pgHost != "" {
		// Individual POSTGRES_* variables from Vercel/Supabase
		DBDSN = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
			pgHost,
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DATABASE"),
			os.Getenv("POSTGRES_PORT"),
		)
	} else {
		// Fallback to original DB_* variables
		DBHost := os.Getenv("DB_HOST")
		DBUser := os.Getenv("DB_USER")
		DBPassword := os.Getenv("DB_PASS")
		DBName := os.Getenv("DB_NAME")
		DBPort := os.Getenv("DB_PORT")

		DBDSN = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s",
			DBHost, DBUser, DBPassword, DBName, DBPort,
		)
	}

	fmt.Println(mylog.ColorizeInfo("\n=========== Setup Database ==========="))
	mylog.Infof("Connecting to database... (DSN Length: %d)", len(DBDSN))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DBDSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		mylog.Errorf("Failed connect to database: %v", err)
		return nil
	}

	mylog.Infof("Success connect to database\n")
	return db
}
