package cmd

import (
	"fmt"

	"os"
	"os/exec"
	"runtime"

	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"github.com/HIMASAKTA-DEV/himasakta-backend/db"
	"github.com/HIMASAKTA-DEV/himasakta-backend/db/migrations"
	dbexport "github.com/HIMASAKTA-DEV/himasakta-backend/scripts/dbexport"
	dbimport "github.com/HIMASAKTA-DEV/himasakta-backend/scripts/dbimport"
	dbclean "github.com/HIMASAKTA-DEV/himasakta-backend/scripts/dbclean"
	seeders "github.com/HIMASAKTA-DEV/himasakta-backend/db/seeder"
	"gorm.io/gorm"
)

func Commands() error {
	db := db.New()
	if err := getParams(db); err != nil {
		return err
	}

	return nil
}

func getParams(db *gorm.DB) error {
	migrate := false
	seeder := false
	watch := false
	test := false
	export := false
	clean := false
	importFile := ""
	storageTarget := ""

	for i := 0; i < len(os.Args[1:]); i++ {
		arg := os.Args[1+i]
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seeder" {
			seeder = true
		}
		if arg == "--watch" {
			watch = true
		}
		if arg == "--test" {
			test = true
		}
		if arg == "--export" {
			export = true
		}
		if arg == "--clean" {
			clean = true
		}
		if arg == "--aws" {
			storageTarget = "s3"
		}
		if arg == "--local" {
			storageTarget = "local"
		}
		if arg == "--import" && i+1 < len(os.Args[1:]) {
			importFile = os.Args[1+i+1]
			i++
		}
	}
	if clean {
		if err := dbclean.CleanData(db); err != nil {
			return fmt.Errorf("clean failed: %w", err)
		}
		mylog.Infof("Database and storage cleaned successfully")
		
		if err := migrations.Migrate(db); err != nil {
			return fmt.Errorf("migration after clean failed: %w", err)
		}
		mylog.Infof("Schema recreated successfully")
	} else if migrate {
		if err := migrations.Migrate(db); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
		mylog.Infof("Migration completed successfully")
	}

	if seeder {
		if db == nil {
			return fmt.Errorf("seeding failed: database connection is nil")
		}
		if err := seeders.Seeding(db); err != nil {
			return fmt.Errorf("seeding failed: %w", err)
		}
		mylog.Infof("Seeder completed successfully")
	}

	if watch {
		if err := runWatch(); err != nil {
			return fmt.Errorf("watching failed: %w", err)
		}
		mylog.Infof("Start watching program")
	}

	if test {
		if err := RunAPITests(); err != nil {
			mylog.Errorf("API tests failed: %v", err)
			os.Exit(1)
		}
	}

	if export {
		if db == nil {
			return fmt.Errorf("export failed: database connection is nil")
		}
		if err := dbexport.ExportDB(db); err != nil {
			return fmt.Errorf("export failed: %w", err)
		}
		mylog.Infof("Export completed successfully")
	}

	if importFile != "" {
		if db == nil {
			return fmt.Errorf("import failed: database connection is nil")
		}
		if err := dbimport.ImportDB(db, importFile, storageTarget); err != nil {
			return fmt.Errorf("import failed: %w", err)
		}
		mylog.Infof("Import completed successfully")
	}

	if clean || seeder || watch || test || export || importFile != "" {
		os.Exit(0)
	}

	return nil
}

func runWatch() error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "air -c .air.windows.toml")
	case "linux", "darwin":
		cmd = exec.Command("bash", "-c", "air -c .air.linux.toml")
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		mylog.Errorf("Error running command: %s", err)
		return err
	}

	mylog.Infoln("Command executed successfully")
	return nil
}
